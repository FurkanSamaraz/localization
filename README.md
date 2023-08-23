# Localization Service Provider Projesi
Bu proje, çoklu dil desteği sağlayan bir hizmet sağlayıcısının temel yapısını içerir. Bu hizmet sağlayıcısı, uygulamaların farklı dillerde metinleri yönetmelerine ve lokalizasyon işlemlerini gerçekleştirmelerine yardımcı olmak için tasarlanmıştır.

## Proje Yapısı
Projede aşağıdaki ana dizinler ve dosyalar bulunmaktadır:

controller: API endpoint'lerini yöneten iş mantığı kodlarının bulunduğu dizin.

controller.go: Temel işlem yönetimi ve endpoint tanımlamalarının yapıldığı dosya.
endpoints: API endpoint'lerinin yönlendirilmesi için kullanılan router'ın bulunduğu dizin.

router.go: Endpoint'lerin belirlendiği ve yönlendirildiği dosya.
locales: Dil çevirilerinin depolandığı dizin.

Latest: En son sürüm çevirilerinin bulunduğu alt dizin.
Denemes: Örnek bir uygulamanın çevirilerinin bulunduğu alt dizin.
v0.json: Dil modülü çeviri verilerinin JSON formatında saklandığı dosya.
storage: Dil çevirilerinin saklandığı depolama mantığının bulunduğu dizin.

storage.go: Çeviri verilerinin okunması, güncellenmesi, oluşturulması ve silinmesi işlemlerinin yapıldığı dosya.
types: Proje genelinde kullanılan veri türlerinin tanımlandığı dizin.

types.go: Veri türü tanımlamalarının yapıldığı dosya.
utils: Yardımcı fonksiyonların ve araçların bulunduğu dizin.

utils.go: Genel amaçlı yardımcı fonksiyonların bulunduğu dosya.
LICENSE: Projenin lisans bilgilerini içeren dosya.

README.md: Proje hakkında genel bilgilerin bulunduğu dosya (şu an okumakta olduğunuz dosya).

go.mod ve go.sum: Projenin bağımlılıklarının yönetildiği Go modül dosyaları.

main.go: Projenin ana başlangıç noktası olan dosya.

## Kullanım
Proje, çoklu dil desteği sağlayan bir hizmet sağlayıcısının temel altyapısını içerdiğinden, farklı uygulamaların farklı dillerdeki çeviri ihtiyaçlarını yönetmek için kullanılabilir. Temel API endpoint'leri aracılığıyla uygulama, dil modülü ve dil bazlı çeviriler oluşturabilir, güncelleyebilir, okuyabilir ve silebilirsiniz.

Projenin çalıştırılması ve test edilmesi hakkında ayrıntılı bilgi için main.go dosyasını inceleyebilirsiniz.

