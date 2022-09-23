Key Word

Day 2
	nested group untuk route bisa dipakai di echo

Day 3
	https redirect dari echo middleware untuk redirect ke https ketika hanya http yang diketik;
	secret key JWT sebaiknya ditaruh di file .env;
	echo.map untuk ubah error status di response;
	di JWT bisa nge-skip endpoint yang tidak ingin digunakan JWT -> pakai Skipper di claimsnya, cek screenshot;

Day 4
	makefile, flag;

Day 7
	github secret;

Day 8 
	Bigbang/replace deployment;
	Rollout deployment;
	BlueGreen deployment;
	Canary deployment;

	tahap deployment => buat server, konek ssh, install docker dan docker compose di server (https://docs.docker.com/engine/install/ubuntu/), sudo docker pull mysql, pull file project dari github, install go, buat RDS di AWS, install mysql di server, setup nginx di server, setting inbound rules di EC2 nya agar bisa diexpose 

Day 9
	Memakai container gunanya untuk menyamakan environment aplikasi saat development dan production;
	Best practice 1 server 1 cotainer 1 image (bisa pakai docker compose), server seharusnya untuk computing sehingga tidak disarankan sebagai tempat penyimpanan data;
	CI adalah proses awal bangun project hingga siap dideploy ke server;
	Golang prosesnya cepat karena tanpa runtime melainkan langsung file executable sehingga langsung mengakses hardware;
	Microservice gunanya untuk menghandle trafic yang berbeda-beda tiap servis, servis dibuat berdasarkan jumlah traficnya, 1 service 1 server 1 DB, komunikasi tiap server bisa pakai Rest-dll;  
	Job/Work Queue;
	Harus bijak dalam penggunaan Cache;
	Database indexing paling sering ditanyakan company, indexing perlu untuk mempercepat pencarian data database;
	HARUS TAU "WHY" KETIKA MENGGUNAKAN SUATU TEKNOLOGI;
	Yang perlu dipelajari DB, Architecture, Algoritma;
