# booking-service

RESTful micro service for my booking application

Create an organization entity that holds a collection of BookableItems and Categories

Create a 'BookableItem' to represent generic service/asset that can be booked for preset times.

```sql
create table category(cat_id int(5) auto_increment primary key, name varchar(50) not null);

create table organization(
    org_id int(5) auto_increment primary key, 
    name varchar(50) not null, 
    website varchar(50),
    phone varchar(20) not null,
    city varchar(30) not null
);
```
This is our join table to manage n-n reln
```sql
create table org_category(orgcat_id int(6) auto_increment primary key,
    org_id int(5) not null,
    cat_id int(5) not null,
    primary key (`org_id`, `cat_id`),
    foreign key (`org_id`) references organization(`org_id`),
    foreign key (`cat_id`) references category(`cat_id`)
);
```
To delete across both tables:
```sql
delete from org_category, organization, category
where organization.org_id=org_category.org_id
and category.cat_id=org_category.cat_id
and category.cat_id=123;
