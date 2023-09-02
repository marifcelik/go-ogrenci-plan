# Go Öğrenci Plan Yönetim Back-End

Bu API öğrenci programlarının yönetileceği bir uygulama varsayılarak geliştirilmiştir.  
Kullanılan teknolojiler:

- Golang
- GORM
- MySQL
- Echo

Kullanıcı giriş işlemleri ve plan yönetimi için **JWT** kullanılmaktadır.

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

Projeyi yalnızca docker ile kullanmak için aşağıdaki komutları çalıştırabilirsiniz.

```bash
# image' ı build edin
docker build -t project-demo:latest .
# db container' ı ayağa kaldırın
docker run --rm -d -p 3306:3306 --network bridge --name project-db -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=project mysql
# server container' ı ayağa kaldırın
docker run --rm -p 3000:3000 --network bridge --name project-demo project-demo:latest
```

Docker Compose kullanarak ayağa kaldırmak için `docker compose up` komutunu kullanmanız yeterli.

## API

| HTTP Metodu | Endpoint      | Eylem         | Açıklama                              |
|-------------|---------------|---------------|---------------------------------------|
| GET         | `/`           | `getAllPlans` | Tüm planları alır.                   |
| GET         | `/:id`        | `getPlan`    | Belirli bir planı alır.              |
| POST        | `/`           | `createPlan` | Yeni bir plan oluşturur.             |
| PATCH       | `/:id`        | `updatePlan` | Belirli bir planı günceller.         |
| DELETE      | `/:id`        | `deletePlan` | Belirli bir planı siler.             |
