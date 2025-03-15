Table user {
    id uuid [primary key]
    firstName string
    lastName string [null]
    email string [null]
    emailVerifiedAt Date [null]
    password string
    createdAt Date
    updatedAt Date
}

Table oauth {
    id uuid [primary key]
    provider "google" | "github" | "facebook"
    identifier "email" | "username"
    metaData json
    userId uuid [ref: > user.id]
    createdAt Date
    updatedAt Date
}

Table file {
    id uuid [primary key]
    type "image" | "video" | "pdf" | "csv" | "excel"
    mimeType string
    fileSize number
    fileName string (example => uuid.jpeg)
    pathName string (example => with out baseurl bucket and folder path)
    description string [null]
    isPrivate boolean
    variants json []
    userId uuid [null, ref: > user.id]
    expenseId uuid [null]
    bankId uuid [null, ref: > bank.id]
    createdAt Date
    updatedAt Date
}

Table bank {
    id uuid [primary key]
    name string
    createdAt Date
    updatedAt Date
}

Table wallet {
    id uuid [primary key]
    name string
    createdAt Date
    updatedAt Date
}

Table userAccount {
    bankId uuid [null, ref: > bank.id]
    walletId uuid [null, ref: > wallet.id]
    userId uuid [ref: > user.id]
    accountNumber string [null]
    phoneNumber string [null]
    name string
    balance number
    createdAt Date
    updatedAt Date
    isActive boolean
}

Table category {
    id uuid [primary key] // 
    name string [unique]
    icon string [null]
    color string
    description string
    createdAt Date
    updatedAt Date
}

Table ledger {
    id uuid [primary key]
    debit number
    credit number
    accountId uuid [null, ref: > userAccount.id]
    transactionId string [null]
    userId uuid [ref: > user.id]
    description string
    date Date
    createdAt Date
    updatedAt Date
}
