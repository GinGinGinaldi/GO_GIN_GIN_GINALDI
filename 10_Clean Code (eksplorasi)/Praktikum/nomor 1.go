package main

import "fmt"

type user struct {
   Email    string
   Password string
}

type userRepo struct {
   DB []user
}

func (r userRepo) Register(u user) {
   if u.Email == "" || u.Password == "" {
       fmt.Println("register failed")
   }

   r.DB = append(r.DB, u)
}

func (r userRepo) Login(u user) string {
   if u.Email == "" || u.Password == "" {
       fmt.Println("login failed")
   }

   for _, us := range r.DB {
       if us.Email == u.Email && us.Password == u.Password {
           return "auth token"
       }
   }

   return ""
}

/* 1. Masalah Return di Fungsi Register
Kalau u.Email atau u.Password kosong, seharusnya kita kasih return setelah nge-print "register failed". 
Soalnya kalau nggak ada return, data user yang invalid tetep bakal ditambahin ke DB. Ini bisa bikin masalah nantinya.

   if u.Email == "" || u.Password == "" {
       fmt.Println("register failed")
       return
   }
2. Masalah di Login (nggak ada return setelah login gagal)
Sama kayak di Register, di Login juga harusnya ada return kalo validasi email atau password kosong. 
Kalo nggak ada return, proses bakal terus jalan walaupun seharusnya gagal. Ini juga bikin nggak efisien.

   if u.Email == "" || u.Password == "" {
       fmt.Println("login failed")
       return ""
   }
3. Pointer Receiver
Di sini userRepo menggunakan value receiver, yang artinya pas kita modif r.DB, 
perubahan nggak akan berdampak ke objek aslinya. Jadi harusnya pakai pointer receiver (*userRepo) biar datanya 
beneran tersimpan.

   func (r *userRepo) Register(u user) {
       r.DB = append(r.DB, u)
   }
4. Error Handling
Penggunaan fmt.Println di dalam fungsi langsung kurang tepat. Harusnya kita balikin error ke caller, 
jadi logikanya bisa lebih fleksibel (misalnya di test atau di aplikasi lain yang bukan CLI).

   if u.Email == "" || u.Password == "" {
       return errors.New("email or password is empty")
   }
5. Hardcoded "auth token" di Login
Hardcode string "auth token" itu nggak aman, sebaiknya nanti diganti sama mekanisme yang 
lebih secure buat nge-generate token. Bisa pake JWT atau semacamnya.
/*