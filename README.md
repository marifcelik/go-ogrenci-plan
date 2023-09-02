# Go Öğrenci Plan Yönetim Back-End
Bu API öğrenci programlarının yönetileceği bir uygulama varsayılarak geliştirilmiştir.  
Kullanılan teknolojiler:
- Golang
- GORM
- MySQL
- Echo

## Kurulum
Direkt olarak localde çalıştıracaksanız MySQL kurulumu yapmanız gerekmektedir. Uygulama DB bağlantısı için `DATABASE_URL` ortam değişkenini kullanmaktadır. 
  
  
Kullanılabilecek ortam değişkenleri:
| Değişken Adı                | Varsayılan Değer | Açıklama                                                        |
|-----------------------------|------------------|-----------------------------------------------------------------|
| `PORT`                      | `3000`           | Uygulamanın çalıştığı port numarası. (Örnek: 3000)              |
| `DATABASE_URL`              | `root:root@tcp(localhost:3306)/project?charset=utf8mb4&parseTime=True&loc=Local` | Veritabanı bağlantı bilgileri, kullanıcı adı, şifre, sunucu vb. |
| `SECRET_KEY`                | `jwt_super_secret_and_super_long_secret_key` | JWT (JSON Web Token) için kullanılan gizli anahtar.             |
| `JWT_EXPIRATION_MINUTES`    | `5` | JWT token'ların geçerlilik süresi dakika cinsinden.             |


MySQL kurulumu yaptıktan sonra aşağıdaki komutları çalıştırarak uygulamayı çalıştırabilirsiniz.
```bash
go mod tidy
go run .
```

### Docker / Docker Compose


