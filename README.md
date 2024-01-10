# GoMailBlast
Golang-based tool to bulk sent emails from csv file using SMTP with rich HTML support with custom reply-to.  


### Installation

1. Clone the repo
```bash
git clone https://github.com/aswinbennyofficial/GoMailBlast.git
```

2. Install dependencies
```bash
go get github.com/joho/godotenv
```

3. Setup environment variables : rename `.env.example` to `.env` and configure it.


4. Customise the message subject and body in the variables `subject` and `body` inside `util/sendEmail.go` file

5. Insert the csv of email list to send in the `data` folder as `data.csv` . Example of `data.csv` is given below : 
```csv
Name,Email
Willy wonka,will@gmail.com
Elon Musk,Elonmusk@gmail.com
Zuckerberg,Zuckerberg@yahoo.com
```

### Usage

```bash
go run .
```