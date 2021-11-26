 create table customer(
    customer_id serial PRIMARY KEY,
    name varchar( 50 ) UNIQUE NOT NULL,
    age INT NOT NULL,
    address varchar(50) NOT NULL
);
