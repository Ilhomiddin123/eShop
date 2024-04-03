INSERT INTO racks (number)
VALUES ('Стелаж A'),
       ('Стелаж Б'),
       ('Стелаж В'),
       ('Стелаж З'),
       ('Стелаж Ж');

INSERT INTO products (title)
VALUES ('Ноутбук'),
       ('Телевизор'),
       ('Телефон'),
       ('Системный блок'),
       ('Часы'),
       ('Микрофон');

INSERT INTO racks_products (rack_id, product_id, priority)
VALUES (1, 2, 'primary'),   -- Телевизор
       (1, 1, 'primary'),   -- Ноутбук
       (2, 3, 'primary'),   -- Телефон
       (3, 3, 'secondary'), -- Телефон (доп.)
       (4, 3, 'secondary'), -- Телефон (доп.)
       (5, 4, 'primary'),   -- Системный блок
       (5, 5, 'primary'),   -- Часы
       (1, 5, 'secondary'), -- Часы (доп.)
       (5, 6, 'secondary'); -- Микрофон


INSERT INTO orders (created_at)
VALUES (current_timestamp),
       (current_timestamp),
       (current_timestamp),
       (current_timestamp),
       (current_timestamp),
       (current_timestamp),
       (current_timestamp),
       (current_timestamp),
       (current_timestamp),
       (current_timestamp),
       (current_timestamp),
       (current_timestamp),
       (current_timestamp),
       (current_timestamp),
       (current_timestamp);

INSERT INTO order_details (order_id, product_id, count)
VALUES (1, 1, 1),
       (2, 1, 1),
       (3, 1, 1),
       (4, 1, 1),
       (5, 1, 1),
       (6, 1, 1),
       (7, 1, 1),
       (8, 1, 1),
       (9, 1, 1),
       (12, 1, 1),
       (13, 1, 1);

INSERT INTO order_details (order_id, product_id, count)
VALUES (10, 1, 2),
       (10, 3, 1),
       (10, 6, 1),
       (11, 2, 3),
       (14, 1, 3),
       (14, 4, 4),
       (15, 5, 1);

SELECT od.order_id, p.title, od.count
FROM order_details od
         JOIN products p on p.id = od.product_id
WHERE order_id in (10, 11, 14, 15)
ORDER BY order_id;

SELECT r.number,
       o.id                           as order_id,
       p.title                        as product_name,
       od.count                       as order_count,
       (SELECT string_agg(r2.number::varchar, ', ')
        FROM racks_products
                 JOIN racks r2 on r2.id = racks_products.rack_id
        WHERE product_id = p.id
          AND priority = 'secondary') as secondary_racks
FROM racks r
         JOIN racks_products rp on r.id = rp.rack_id
         JOIN products p on p.id = rp.product_id
         JOIN order_details od on p.id = od.product_id
         JOIN orders o on o.id = od.order_id
WHERE r.id = 5
  AND o.id in (10, 11, 14, 15)
  AND rp.priority = 'primary';


SELECT string_agg(rack_id::varchar, ', ')
FROM racks_products
WHERE product_id = 3
  AND priority = 'secondary';