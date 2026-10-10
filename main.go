package main

import "fmt"

// Notifier la interface ma ung dung su dung de gui thong bao.
type Notifier interface {
	Send(message string) error
}

// LegacySMS la thu vien cũ. Khong chỉnh sửa struct va phuong thuc nay.
type LegacySMS struct{}

func (s *LegacySMS) SendSMS(content string) {
	fmt.Println("SMS:", content)
}

// SMSAdapter là adapter kết nối LegacySMS với interface Notifier.
// Nó chứa một con trỏ tới LegacySMS và "chuyển đổi" phương thức SendSMS
// sang phương thức Send mà interface Notifier yêu cầu.
type SMSAdapter struct {
	sms *LegacySMS
}

// Send triển khai interface Notifier.
// Phương thức này gọi LegacySMS.SendSMS() để gửi tin nhắn,
// đồng thời trả về nil vì LegacySMS không phát sinh lỗi.
func (a *SMSAdapter) Send(message string) error {
	a.sms.SendSMS(message)
	return nil
}

func main() {
	// Khởi tạo thư viện cũ LegacySMS
	legacy := &LegacySMS{}

	// Tạo SMSAdapter bọc thư viện cũ
	adapter := &SMSAdapter{sms: legacy}

	// Gán adapter cho biến kiểu Notifier (interface)
	var notifier Notifier = adapter

	// Gọi Send thông qua interface Notifier
	notifier.Send("Xin chao, don hang cua ban da duoc xac nhan!")
}
