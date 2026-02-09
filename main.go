package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const (
	defaultTitle = "喝水提醒"
	defaultDesp  = "该喝水啦 💧"
)

func main() {
	uid := strings.TrimSpace(os.Getenv("PUSH_UID"))
	if uid == "" {
		exitWithError(errors.New("缺少 uid：请设置环境变量 PUSH_UID"))
	}

	sendKey := strings.TrimSpace(os.Getenv("SENDKEY"))
	if sendKey == "" {
		exitWithError(errors.New("缺少 sendkey：请设置环境变量 SENDKEY（建议来自 GitHub repository secret）"))
	}

	if err := sendReminder(uid, sendKey, defaultTitle, defaultDesp); err != nil {
		exitWithError(err)
	}

	fmt.Println("发送成功：喝水提醒已推送")
}

func sendReminder(uid, sendKey, title, desp string) error {
	endpoint := fmt.Sprintf("https://%s.push.ft07.com/send/%s.send", uid, sendKey)

	q := url.Values{}
	q.Set("title", title)
	q.Set("desp", desp)

	requestURL := endpoint + "?" + q.Encode()

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(requestURL)
	if err != nil {
		return fmt.Errorf("请求失败: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("发送失败，HTTP %d: %s", resp.StatusCode, strings.TrimSpace(string(body)))
	}

	fmt.Printf("服务响应: %s\n", strings.TrimSpace(string(body)))
	return nil
}

func exitWithError(err error) {
	fmt.Fprintln(os.Stderr, "错误:", err)
	os.Exit(1)
}
