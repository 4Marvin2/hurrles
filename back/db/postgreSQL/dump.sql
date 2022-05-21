create extension if not exists citext;

create table if not exists users
(
    id          serial primary key,
    email       citext unique,
    password    varchar(64) not null,
    full_name   varchar(128) default '',
    number      varchar(16) default '',
    is_admin    boolean default false
);

create table if not exists restaurants
(
    id          serial primary key,
    title       varchar(255) not null,
    description varchar(1000) not null,
    address     varchar(255) not null,
    metro       varchar(50) not null,
    number      varchar(16)  not null,
    open_time   timestamptz  not null,
    close_time  timestamptz  not null,
    kitchen     varchar(255) not null,
    img         varchar(255) not null
);

create table if not exists restaurants_users
(
    id       serial primary key,
    user_id int not null,
    constraint fk_ru_users foreign key (user_id) references users (id),
    restaurant_id  int not null,
    constraint fk_ru_restaurants foreign key (restaurant_id) references restaurants (id),
    unique(user_id, restaurant_id)
);

create table if not exists places
(
    id            serial primary key,
    restaurant_id int not null,
    constraint fk_dishes_restaurants foreign key (restaurant_id) references restaurants (id),
    capacity      int not null,
    number        int not null,
    left_top      double precision not null,
    right_bottom  double precision not null,
    floor         int not null
);

create table if not exists orders
(
    id           serial primary key,
    user_id      int         not null,
    constraint fk_orders_users foreign key (user_id) references users (id),
    place_id     int not null,
    constraint fk_orders_places foreign key (place_id) references places (id),
    start_time   timestamptz not null,
    end_time     timestamptz not null,
    cost         int         not null,
    created_time timestamptz default now()
    status       varchar(15) not null default 'open'
);

create table if not exists payments
(
    id         serial primary key,
    order_id   int not null,
    constraint fk_payments_orders foreign key (order_id) references orders (id),
    total_cost int not null,
    format     int not null,
    status     int not null
);

create table if not exists dishes
(
    id            serial primary key,
    restaurant_id int          not null,
    constraint fk_dishes_restaurants foreign key (restaurant_id) references restaurants (id),
    title         varchar(255) not null,
    description   varchar(255) not null,
    price         int          not null
);

create table if not exists dishes_orders
(
    id       serial primary key,
    order_id int not null,
    constraint fk_do_orders foreign key (order_id) references orders (id),
    dish_id  int not null,
    constraint fk_do_dishes foreign key (dish_id) references dishes (id),
    number   int not null default 1
);
