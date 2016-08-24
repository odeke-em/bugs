#!/usr/bin/env python3

data = ",Carl,Gauss,cgauss@unigottingen.de,Male,30.4.17.77\n"
with open("mock_data.csv", "w") as f:
    f.write("id,first_name,last_name,email,gender,ip_address\n")
    f.write(("1"+data)*int(1e6))
    f.write("42"+data);
