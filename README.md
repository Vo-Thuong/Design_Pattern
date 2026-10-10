# Bài thi: Adapter Pattern

```text
Họ và tên: Võ Công Thương
Mã sinh viên: 145010124001
Lớp: SE-K24
Bài thi: Adapter Pattern
```

## Hướng dẫn chạy chương trình

Từ thư mục chứa `go.mod` (Go 1.22 trở lên):

```bash
go build ./...
go run .
```

Kết quả khi chạy thành công:

```text
SMS: Xin chao, don hang cua ban da duoc xac nhan!
```

## Câu trả lời của sinh viên

1. **Adapter Pattern giải quyết vấn đề gì?**

   Adapter Pattern giải quyết vấn đề tương thích giữa các interface không khớp nhau. Khi ứng dụng yêu cầu một interface cụ thể (ở đây là `Notifier` với phương thức `Send`), nhưng thư viện có sẵn lại dùng interface khác (`LegacySMS` với phương thức `SendSMS`), Adapter đóng vai trò cầu nối, chuyển đổi lời gọi từ interface mới sang interface cũ mà không cần thay đổi bất kỳ bên nào.

2. **Tại sao không cần chỉnh sửa `LegacySMS`?**

   Vì mục tiêu của Adapter Pattern chính là tích hợp các thành phần cũ (thư viện bên thứ ba, legacy code) vào hệ thống mới mà **không** thay đổi mã nguồn gốc. `SMSAdapter` đảm nhận toàn bộ trách nhiệm chuyển đổi: nó nhận lời gọi `Send()` từ phía ứng dụng rồi uỷ quyền (delegate) sang `SendSMS()` của `LegacySMS`. Nhờ vậy, nguyên tắc Open/Closed được tuân thủ — mở rộng hành vi mà không cần sửa code đã có.

## Nộp bài

Nộp `main.go`, `go.mod`, `README.md` và link GitHub repository. Điền đủ thông tin cá nhân phía trên và bảo đảm người chấm truy cập được repository.

