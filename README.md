# Deprem İzleme Sistemi

Bu proje, gerçek zamanlı olarak deprem verilerini toplar, işler ve anormal depremleri bir harita üzerinde gösterir.

## Başlangıç

Bu projeyi yerel geliştirme ortamınızda çalıştırmak için aşağıdaki adımları izleyin.

### Önkoşullar

- Docker

### Kurulum

#### 1. Projeyi yerel makinenize klonlayın:

```bash
git clone hhttps://github.com/ouzzkp/earthquake-app
cd earthquake-app
```
#### Docker ile projeyi başlatın:
```bash
docker-compose up --build
```
 Bu komut, gereken tüm servisleri Docker container'larında başlatır. Backend servisi localhost:8080 üzerinde çalışır.


## API Kullanımı

#### Tüm Depremleri Listele

```http
  GET /api/earthquakes

```

| Parametre | Tip     | Açıklama                |
| :-------- | :------- | :------------------------- |
| `api_key` | `string` | **Gerekli**. API anahtarınız. |

#### Tek Bir Depremi getir

```http
  GET /api/earthquakes/{id}
```

| Parametre | Tip     | Açıklama                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `int` | **Gerekli**. Depremin benzersiz ID'si. | 
| `api_key`      | `string` | **Gerekli**. API anahtarınız. |

#### Yeni Bir Deprem Ekleme

```http
  POST /api/earthquakes
```

| Parametre | Tip     | Açıklama                       |
| :-------- | :------- | :-------------------------------- |
| `Latitude`      | `string` | **Gerekli**. Depremin enlemi. | 
| `Longitude`      | `string` | **Gerekli**. Depremin boylamı. |
| `Magnitude`      | `float` | **Gerekli**.  Depremin şiddeti. | 
| `Time`      | `string` | **Gerekli**. Depremin zamanı. |
| `api_key`      | `string` | **Gerekli**. API anahtarınız. |

#### Deprem Verisini Güncelle

```http
  PUT /api/earthquakes/{id}
```

| Parametre | Tip     | Açıklama                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `int` | **Gerekli**. Depremin benzersiz ID'si. | 
| `api_key`      | `string` | **Gerekli**. API anahtarınız. |

#### Deprem Verisini Silme

```http
  DELETE /api/earthquakes/{id}
```

| Parametre | Tip     | Açıklama                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `int` | **Gerekli**. Depremin benzersiz ID'si. | 
| `api_key`      | `string` | **Gerekli**. API anahtarınız. |



## Kullanılan Teknolojiler

**Backend:**

- Go (Golang): Sunucu tarafı mantığını oluşturma, HTTP isteklerini işleme, veritabanı ile etkileşim ve gerçek zamanlı veri aktarımı için WebSocket kullanımı.
- PostgreSQL: Deprem verilerini depolamak ve yönetmek için tercih edilen ilişkisel veritabanı.
- WebSocket: Sunucu ile istemci arasında gerçek zamanlı iletişim sağlamak, özellikle yeni deprem bilgilerinin anında yayınlanması için kullanılır.


**Frontend:** 

- React: Kullanıcı arayüzünü oluşturmak için kullanılan bir JavaScript kütüphanesi, özellikle haritada deprem verilerini gösterme ve bu bilgileri gerçek zamanlı olarak güncelleme.
- react-simple-maps: Haritayı render etmek ve üzerine deprem işaretçileri yerleştirmek için kullanılan React bileşen kütüphanesi.

  
