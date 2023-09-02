# Go Öğrenci Plan Yönetim Back-End

Bu API öğrenci programlarının yönetileceği bir uygulama varsayılarak geliştirilmiştir.  
Kullanılan teknolojiler:

- Golang
- GORM
- MySQL
- Echo

Kullanıcı giriş işlemleri ve plan yönetimi için **JWT** kullanılmaktadır.

## Özellikler
- Belirli tarih ve saat aralıklarında plan oluşturma
- Planları güncelleme, silme
- Planları haftalık ve aylık olarak görüntüleme
- Planlar üzerinde durum belirleme (Yapılıyor, Bitti, İptal)
- Aynı aralıkta birden fazla plan olması halinde uyarı verme
- Kullanıcı kayıt, giriş ve güncelleme işlemleri

## Kurulum

Direkt olarak localde çalıştıracaksanız MySQL kurulumu yapmanız gerekmektedir. Uygulama DB bağlantısı için `DATABASE_URL` ortam değişkenini kullanmaktadır.
  
Kullanılabilecek ortam değişkenleri:
| Değişken Adı             | Varsayılan Değer                             | Açıklama                                                        |
| ------------------------ | -------------------------------------------- | --------------------------------------------------------------- |
| `PORT`                   | `3000`                                       | Uygulamanın çalıştığı port numarası. (Örnek: 3000)              |
| `DATABASE_URL`           | `root:root@tcp(localhost:3306)/project`      | Veritabanı bağlantı bilgileri, kullanıcı adı, şifre, sunucu vb. |
| `SECRET_KEY`             | `jwt_super_secret_and_super_long_secret_key` | JWT (JSON Web Token) için kullanılan gizli anahtar.             |
| `JWT_EXPIRATION_MINUTES` | `5`                                          | JWT token'ların geçerlilik süresi dakika cinsinden.             |

MySQL kurulumu yaptıktan sonra aşağıdaki komutları çalıştırarak uygulamayı çalıştırabilirsiniz.

```bash
go mod tidy
go run .
```

### Docker Compose

Docker Compose kullanarak ayağa kaldırmak için `docker compose up` komutunu kullanmanız yeterli. Bu komut 2 adet container ayağa kaldıracaktır. İlki kök dizinde bulunan Dockerfile' ı kullanarak uygulamayı ayağa kaldıracaktır. İkincisi ise  MySQL imajını kullanarak bir veritabanı ayağa kaldıracaktır.

## API

### /plan  
| HTTP Metodu | Endpoint | Girdiler                       | Açıklama                          |
| ----------- | -------- | ------------------------------ | --------------------------------- |
| GET         | `/`      | -                              | Tüm planları getirir.             |
| GET         | `/:id`   | -                              | Verilen id' ye ait planı getirir. |
| POST        | `/`      | `title, content, start?, end?` | Yeni bir plan oluşturur.          |
| PATCH       | `/:id`   | `title, content, start?, end?` | Belirli bir planı günceller.      |
| DELETE      | `/:id`   | -                              | Belirli bir planı siler.          |
  
**plan yapısı**
```json
{
  "id": 1,
  "title": "Plan Başlığı",
  "content": "Plan Açıklaması",
  "start": "2021-01-01T00:00:00Z",
  "end": "2021-01-01T00:00:00Z",
  "state": 0,
  "created_at": "2021-01-01T00:00:00Z",
  "updated_at": "2021-01-01T00:00:00Z",
  "student_id": 1
}
```

### /auth
| HTTP Metodu | Endpoint   | Girdiler                            | Açıklama                          |
| ----------- | ---------- | ----------------------------------- | --------------------------------- |
| POST        | `/signin`  | `username, password`                | Kullanıcı girişini işler.         |
| POST        | `/signup`  | `name, surname, username, password` | Yeni bir kullanıcı kaydını işler. |
| POST        | `/signout` | -                                   | Kullanıcı çıkışını işler.         |

**user yapısı**
```json
{
    "id": 1,
    "name": "Ad",
    "surname": "Soyad",
    "username": "kullanici_adi",
    "password": "sifre"
}
```

### /profile
| HTTP Metodu | Endpoint | Girdiler                     | Açıklama                                                |
| ----------- | -------- | ---------------------------- | ------------------------------------------------------- |
| GET         | `/`      | -                            | Giriş yapmış kullanıcının profil bilgilerini alır.      |
| PATCH       | `/`      | `name?, surname?, password?` | Giriş yapmış kullanıcının profil bilgilerini günceller. |

