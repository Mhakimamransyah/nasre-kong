# Simulasi Multi Services menggunakan Kong Gateway

## Workaround
- Buat file logging (gateway.log) di dalam folder logs (plugin log file)
- docker compose up
- Buat koneksi dari konga ke kong dengan menggunakan nama kontainer atau ip pada jaringan yang sama , contoh :
   - http://kong:8001 (akses ke kontainer kong melalui internal network)
   - http://192.168.102.103:8001 (akses ke kontainer kong melalui internet)
## Konsep
- 1 servis = 1 host = n route
- bisa menggunakan consumer untuk apply rule ke spesific route atau services

## Authentikasi
- JWT atau API KEY
- API KEY (set api key di spesific route ,atur melalui consumer)
- JWT (set jwt iss dan secret)

## Logging
- masih pakai log file (butuh riset lebih lanjut)

## Proxy Cache
https://docs.konghq.com/gateway/latest/get-started/proxy-caching/

## Rate Limiting
https://docs.konghq.com/gateway/latest/get-started/rate-limiting/

## Load Balancer
https://docs.konghq.com/gateway/latest/get-started/load-balancing/ 

**CATATAN :** untuk menguji load balancer matikan plugin cache terlebih dahulu karena kalo cache hidup maka kong akan selalu menggunakan cache dan tidak menggunakan servis lainnya