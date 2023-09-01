# Golang, chi ve MySQL ile Ürün CRUD API Oluşturma

Bu doküman, Golang ile `chi` kütüphanesini kullanarak ve MySQL veritabanı ile birlikte temel bir ürün CRUD API'si oluşturmak için bir örneği içermektedir.

## Gerekli Paketleri Kurma:

Projede kullanılan paketleri yüklemek için aşağıdaki komutları çalıştırabilirsiniz:

```bash
go get -u github.com/go-chi/chi
go get -u github.com/go-sql-driver/mysql
go get -u github.com/jmoiron/sqlx
```



###  Golang ile chi kütüphanesinin özellikleri?

`chi` Golang için modern bir HTTP yönlendirici (router) kütüphanesidir. Go'nun standart `net/http` paketine dayanarak, etkili ve ekspresif bir şekilde HTTP yönlendirmeleri oluşturmanıza olanak tanır.

`chi` kütüphanesinin bazı özellikleri:
1. **Orta katman desteği**: `chi` middleware desteği ile gelir. Bu, belirli bir HTTP talebine önce veya sonra işlem yapabilmeniz için basit ve etkili bir yöntem sunar.
2. **Parametreli URL Yönlendirme**: `chi` URL parametrelerini destekler. Örneğin, `/users/{userID}` şeklinde bir yönlendirme tanımlayarak, `userID` parametresini dinamik olarak alabilirsiniz.
3. **Yerleşik Middleware**: `chi` bazı yaygın middleware'leri içerir. Bu, özellikle loglama, sıkıştırma, önbellekleme vb. işlevleri için kullanışlıdır.
4. **Hafif ve Hızlı**: `chi` minimalist ve hızlıdır. Gereksiz işlevsellik eklenmez, bu nedenle uygulamanızın performansını en üst düzeyde tutmaya yardımcı olur.
5. **Mux**'tan İlham Aldı**: `chi`, Go'nun standart `http.ServeMux`'ından ilham alır ve onunla uyumludur, bu da kodunuzu taşımak ve mevcut Go HTTP koduyla entegre olmak için idealdir.

`chi` kullanarak bir HTTP sunucusu oluşturmak oldukça basittir:


### jmoiron/sqlx kütüphanesinin özellikleri?
`github.com/jmoiron/sqlx` kütüphanesi, Go'nun standart `database/sql` kütüphanesine genel bir uzantıdır. `sqlx` bazı ekstra özellikler ve kolaylıklar sağlar, böylece Go programcıları veritabanı işlemlerini daha verimli ve hatasız bir şekilde gerçekleştirebilirler.

`sqlx`'in bazı öne çıkan özellikleri şunlardır:

1. **Daha Zengin Sorgular**: İsme dayalı bağlama (`:name` gibi) ile sorgu oluşturma kolaylaşır.
2. **Yapılarla Doğrudan Bağlama**: Sonuçları doğrudan Go yapısına bağlama imkanı sunar, bu da manuel alan atamalarını azaltır.
3. **Genişletilmiş Sorgu Fonksiyonları**: `Get()`, `Select()`, `NamedExec()` gibi ekstra sorgu fonksiyonları sunar.
4. **İlişkisel Veritabanı Desteği**: Hem SQL hem de NoSQL veritabanları için sürücüleri destekler.
5. **Daha Kullanıcı Dostu API**: Go'nun standart `database/sql` kütüphanesi ile bire bir uyumlu olsa da, daha kullanıcı dostu bir API'ye sahiptir.

Kısacası, `sqlx`, Go'da veritabanı işlemleri yaparken işleri daha kolay ve hızlı hale getiren genişletilmiş bir arayüzdür. Eğer bir Go uygulamasında veritabanı işlemleri yapıyorsanız, `sqlx`'i kullanarak işlerinizi hızlandırabilir ve kodunuzu daha okunaklı hale getirebilirsiniz.