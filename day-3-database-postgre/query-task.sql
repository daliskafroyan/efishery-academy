
drop table if exists customers cascade;
drop table if exists products cascade;
drop table if exists orders cascade;

create table if not exists customers (
	id serial primary key, 
	customer_name char(50) not null
);

create table if not exists products (
	id serial primary key, 
	name char(50) not null
);

create table if not exists orders (
	order_id serial primary key, 
	customer_id int references customers (id),
	product_id int references products (id),
	order_date date not null,
	total float not null
);


-- insert new data

insert into products (name) values ('Pakan Udang');
insert into products (name) values ('Pakan Kucing');
insert into products (name) values ('Pakan Nila');

insert into customers (customer_name) values ('Yae');
insert into customers (customer_name) values ('Kaedehara');
insert into customers (customer_name) values ('Kazuha');

insert into orders (customer_id, product_id, order_date, total) values (1,1, '12-10-2020', 100);
insert into orders (customer_id, product_id, order_date, total) values (2,2, '2-10-2020',70000);
insert into orders (customer_id, product_id, order_date, total) values (3,3, '12-9-2020',30000);


-- update data

update customers set customer_name = 'Inazuma' where customer_name = 'Yae';
update products set name = 'Pakan Burung' where name = 'Pakan Udang';
update orders set total = 999999 where customer_id = 1;


-- delete data

delete from orders where order_date = '2-10-2020';
delete from products where name = 'Pakan Kucing';
delete from customers where customer_name = 'Kaedehara';


-- join data

select customers.customer_name, products.name, orders.order_date, orders.total
from customers inner join orders on orders.customer_id = customers.id
inner join products on orders.product_id =  products.id