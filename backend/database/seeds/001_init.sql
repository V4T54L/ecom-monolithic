INSERT INTO
    users (
        name,
        username,
        email,
        password,
        role,
        created_at,
        updated_at
    )
VALUES (
        'Alice Smith',
        'alice.smith',
        'alice.smith@example.com',
        'password123',
        'user',
        NOW(),
        NOW()
    ),
    (
        'Bob Johnson',
        'bob.johnson',
        'bob.johnson@example.com',
        'password123',
        'user',
        NOW(),
        NOW()
    ),
    (
        'Charlie Brown',
        'charlie.brown',
        'charlie.brown@example.com',
        'password123',
        'admin',
        NOW(),
        NOW()
    ),
    (
        'David Wilson',
        'david.wilson',
        'david.wilson@example.com',
        'password123',
        'user',
        NOW(),
        NOW()
    ),
    (
        'Eve Davis',
        'eve.davis',
        'eve.davis@example.com',
        'password123',
        'user',
        NOW(),
        NOW()
    ),
    (
        'Frank Miller',
        'frank.miller',
        'frank.miller@example.com',
        'password123',
        'user',
        NOW(),
        NOW()
    ),
    (
        'Grace Lee',
        'grace.lee',
        'grace.lee@example.com',
        'password123',
        'user',
        NOW(),
        NOW()
    ),
    (
        'Hank Taylor',
        'hank.taylor',
        'hank.taylor@example.com',
        'password123',
        'user',
        NOW(),
        NOW()
    ),
    (
        'Ivy Thomas',
        'ivy.thomas',
        'ivy.thomas@example.com',
        'password123',
        'user',
        NOW(),
        NOW()
    ),
    (
        'Jack Jackson',
        'jack.jackson',
        'jack.jackson@example.com',
        'password123',
        'user',
        NOW(),
        NOW()
    ),
    (
        'Katie White',
        'katie.white',
        'katie.white@example.com',
        'password123',
        'user',
        NOW(),
        NOW()
    ),
    (
        'Liam Harris',
        'liam.harris',
        'liam.harris@example.com',
        'password123',
        'user',
        NOW(),
        NOW()
    ),
    (
        'Mia Clark',
        'mia.clark',
        'mia.clark@example.com',
        'password123',
        'user',
        NOW(),
        NOW()
    ),
    (
        'Noah Lewis',
        'noah.lewis',
        'noah.lewis@example.com',
        'password123',
        'user',
        NOW(),
        NOW()
    ),
    (
        'Olivia Walker',
        'olivia.walker',
        'olivia.walker@example.com',
        'password123',
        'user',
        NOW(),
        NOW()
    ),
    (
        'Paul Hall',
        'paul.hall',
        'paul.hall@example.com',
        'password123',
        'user',
        NOW(),
        NOW()
    ),
    (
        'Quincy Allen',
        'quincy.allen',
        'quincy.allen@example.com',
        'password123',
        'user',
        NOW(),
        NOW()
    ),
    (
        'Rachel Young',
        'rachel.young',
        'rachel.young@example.com',
        'password123',
        'user',
        NOW(),
        NOW()
    ),
    (
        'Steve Hernandez',
        'steve.hernandez',
        'steve.hernandez@example.com',
        'password123',
        'user',
        NOW(),
        NOW()
    ),
    (
        'Tina King',
        'tina.king',
        'tina.king@example.com',
        'password123',
        'user',
        NOW(),
        NOW()
    );

INSERT INTO
    addresses (
        user_id,
        address,
        city,
        zipcode
    )
VALUES (
        1,
        '123 Maple St',
        'Springfield',
        '12345'
    ),
    (
        2,
        '456 Oak St',
        'Springfield',
        '12346'
    ),
    (
        3,
        '789 Pine St',
        'Shelbyville',
        '12347'
    ),
    (
        4,
        '321 Birch St',
        'Springfield',
        '12348'
    ),
    (
        5,
        '654 Cedar St',
        'Shelbyville',
        '12349'
    ),
    (
        6,
        '987 Elm St',
        'Capitol City',
        '12350'
    ),
    (
        7,
        '159 Spruce St',
        'Springfield',
        '12351'
    ),
    (
        8,
        '753 Willow St',
        'Springfield',
        '12352'
    ),
    (
        9,
        '852 Chestnut St',
        'Shelbyville',
        '12353'
    ),
    (
        10,
        '963 Ash St',
        'Capitol City',
        '12354'
    ),
    (
        11,
        '741 Maple Ave',
        'Springfield',
        '12355'
    ),
    (
        12,
        '258 Oak Ave',
        'Springfield',
        '12356'
    ),
    (
        13,
        '369 Pine Ave',
        'Shelbyville',
        '12357'
    ),
    (
        14,
        '147 Birch Ave',
        'Springfield',
        '12358'
    ),
    (
        15,
        '258 Cedar Ave',
        'Shelbyville',
        '12359'
    ),
    (
        16,
        '369 Elm Ave',
        'Capitol City',
        '12360'
    ),
    (
        17,
        '456 Spruce Ave',
        'Springfield',
        '12361'
    ),
    (
        18,
        '537 Willow Ave',
        'Springfield',
        '12362'
    ),
    (
        19,
        '648 Chestnut Ave',
        'Shelbyville',
        '12363'
    ),
    (
        20,
        '759 Ash Ave',
        'Capitol City',
        '12364'
    );

INSERT INTO
    categories (name, start, end)
VALUES (
        'Electronics',
        '2023-01-01',
        '2023-12-31'
    ),
    (
        'Books',
        '2023-01-01',
        '2023-12-31'
    ),
    (
        'Clothing',
        '2023-01-01',
        '2023-12-31'
    ),
    (
        'Furniture',
        '2023-01-01',
        '2023-12-31'
    ),
    (
        'Toys',
        '2023-01-01',
        '2023-12-31'
    ),
    (
        'Sports',
        '2023-01-01',
        '2023-12-31'
    ),
    (
        'Beauty',
        '2023-01-01',
        '2023-12-31'
    ),
    (
        'Grocery',
        '2023-01-01',
        '2023-12-31'
    ),
    (
        'Automotive',
        '2023-01-01',
        '2023-12-31'
    ),
    (
        'Home Improvement',
        '2023-01-01',
        '2023-12-31'
    ),
    (
        'Pet Supplies',
        '2023-01-01',
        '2023-12-31'
    ),
    (
        'Jewelry',
        '2023-01-01',
        '2023-12-31'
    ),
    (
        'Health',
        '2023-01-01',
        '2023-12-31'
    ),
    (
        'Garden',
        '2023-01-01',
        '2023-12-31'
    ),
    (
        'Travel',
        '2023-01-01',
        '2023-12-31'
    ),
    (
        'Tools',
        '2023-01-01',
        '2023-12-31'
    ),
    (
        'Music',
        '2023-01-01',
        '2023-12-31'
    ),
    (
        'Software',
        '2023-01-01',
        '2023-12-31'
    ),
    (
        'Crafts',
        '2023-01-01',
        '2023-12-31'
    ),
    (
        'Office Supplies',
        '2023-01-01',
        '2023-12-31'
    );

INSERT INTO
    products (
        thumbnail_url,
        title,
        description,
        rating,
        price,
        discountedPrice,
        category_id,
        stock,
        created_at,
        updated_at
    )
VALUES (
        'https://placehold.co/600x400?text=img1.jpg',
        'Smartphone',
        'Latest model smartphone with all features.',
        4.5,
        699.99,
        649.99,
        1,
        100,
        NOW(),
        NOW()
    ),
    (
        'https://placehold.co/600x400?text=img2.jpg',
        'Laptop',
        'High-performance laptop for all your needs.',
        4.7,
        1299.99,
        1249.99,
        1,
        50,
        NOW(),
        NOW()
    ),
    (
        'https://placehold.co/600x400?text=img3.jpg',
        'Headphones',
        'Noise-canceling over-ear headphones.',
        4.2,
        199.99,
        179.99,
        1,
        200,
        NOW(),
        NOW()
    ),
    (
        'https://placehold.co/600x400?text=img4.jpg',
        'Novel Book',
        'An exciting novel for all book lovers.',
        4.8,
        15.99,
        12.99,
        2,
        300,
        NOW(),
        NOW()
    ),
    (
        'https://placehold.co/600x400?text=img5.jpg',
        'History Book',
        'A comprehensive history of the world.',
        4.5,
        20.99,
        17.99,
        2,
        250,
        NOW(),
        NOW()
    ),
    (
        'https://placehold.co/600x400?text=img6.jpg',
        'T-shirt',
        'Comfortable cotton T-shirt in various sizes.',
        4.0,
        24.99,
        19.99,
        3,
        150,
        NOW(),
        NOW()
    ),
    (
        'https://placehold.co/600x400?text=img7.jpg',
        'Jeans',
        'Stylish denim jeans for everyday wear.',
        4.3,
        49.99,
        39.99,
        3,
        120,
        NOW(),
        NOW()
    ),
    (
        'https://placehold.co/600x400?text=img8.jpg',
        'Sofa',
        'Luxurious sofa for living room comfort.',
        4.9,
        799.99,
        749.99,
        4,
        30,
        NOW(),
        NOW()
    ),
    (
        'https://placehold.co/600x400?text=img9.jpg',
        'Coffee Table',
        'Elegant coffee table for your living room.',
        4.6,
        299.99,
        249.99,
        4,
        50,
        NOW(),
        NOW()
    ),
    (
        'https://placehold.co/600x400?text=img10.jpg',
        'Action Figure',
        'Collectible action figure from a popular franchise.',
        4.5,
        29.99,
        24.99,
        5,
        300,
        NOW(),
        NOW()
    ),
    (
        'https://placehold.co/600x400?text=img11.jpg',
        'Soccer Ball',
        'High-quality soccer ball for matches.',
        4.2,
        39.99,
        34.99,
        6,
        150,
        NOW(),
        NOW()
    ),
    (
        'https://placehold.co/600x400?text=img12.jpg',
        'Makeup Kit',
        'Complete makeup kit for beauty enthusiasts.',
        4.8,
        59.99,
        49.99,
        7,
        70,
        NOW(),
        NOW()
    ),
    (
        'https://placehold.co/600x400?text=img13.jpg',
        'Organic Apples',
        'Fresh organic apples from local farms.',
        4.7,
        3.99,
        2.99,
        8,
        500,
        NOW(),
        NOW()
    ),
    (
        'https://placehold.co/600x400?text=img14.jpg',
        'Car Oil',
        'High-quality synthetic oil for vehicles.',
        4.5,
        29.99,
        24.99,
        9,
        80,
        NOW(),
        NOW()
    ),
    (
        'https://placehold.co/600x400?text=img15.jpg',
        'Drill Set',
        'Set of power drills for construction projects.',
        4.6,
        99.99,
        89.99,
        10,
        60,
        NOW(),
        NOW()
    ),
    (
        'https://placehold.co/600x400?text=img16.jpg',
        'Dog Collar',
        'Comfortable collar for your furry friends.',
        4.4,
        15.99,
        12.99,
        11,
        200,
        NOW(),
        NOW()
    ),
    (
        'https://placehold.co/600x400?text=img17.jpg',
        'Gold Ring',
        'Elegant gold ring for special occasions.',
        4.9,
        499.99,
        479.99,
        12,
        25,
        NOW(),
        NOW()
    ),
    (
        'https://placehold.co/600x400?text=img18.jpg',
        'Herbal Tea',
        'Natural herbal tea for a refreshing drink.',
        4.8,
        9.99,
        7.99,
        13,
        300,
        NOW(),
        NOW()
    ),
    (
        'https://placehold.co/600x400?text=img19.jpg',
        'Yoga Mat',
        'Comfortable mat for yoga and exercise.',
        4.5,
        25.99,
        19.99,
        14,
        150,
        NOW(),
        NOW()
    ),
    (
        'https://placehold.co/600x400?text=img20.jpg',
        'Home Tool Set',
        'Complete tool set for home improvement.',
        4.6,
        79.99,
        69.99,
        15,
        40,
        NOW(),
        NOW()
    );

INSERT INTO
    product_images (product_id, image_url)
VALUES (
        1,
        'https://placehold.co/600x400?text=img1-a.jpg'
    ),
    (
        1,
        'https://placehold.co/600x400?text=img1-b.jpg'
    ),
    (
        2,
        'https://placehold.co/600x400?text=img2-a.jpg'
    ),
    (
        2,
        'https://placehold.co/600x400?text=img2-b.jpg'
    ),
    (
        3,
        'https://placehold.co/600x400?text=img3-a.jpg'
    ),
    (
        4,
        'https://placehold.co/600x400?text=img4-a.jpg'
    ),
    (
        4,
        'https://placehold.co/600x400?text=img4-b.jpg'
    ),
    (
        5,
        'https://placehold.co/600x400?text=img5-a.jpg'
    ),
    (
        6,
        'https://placehold.co/600x400?text=img6-a.jpg'
    ),
    (
        7,
        'https://placehold.co/600x400?text=img7-a.jpg'
    ),
    (
        8,
        'https://placehold.co/600x400?text=img8-a.jpg'
    ),
    (
        9,
        'https://placehold.co/600x400?text=img9-a.jpg'
    ),
    (
        10,
        'https://placehold.co/600x400?text=img10-a.jpg'
    ),
    (
        11,
        'https://placehold.co/600x400?text=img11-a.jpg'
    ),
    (
        12,
        'https://placehold.co/600x400?text=img12-a.jpg'
    ),
    (
        13,
        'https://placehold.co/600x400?text=img13-a.jpg'
    ),
    (
        14,
        'https://placehold.co/600x400?text=img14-a.jpg'
    ),
    (
        15,
        'https://placehold.co/600x400?text=img15-a.jpg'
    ),
    (
        16,
        'https://placehold.co/600x400?text=img16-a.jpg'
    ),
    (
        17,
        'https://placehold.co/600x400?text=img17-a.jpg'
    ),
    (
        18,
        'https://placehold.co/600x400?text=img18-a.jpg'
    );

INSERT INTO product_reviews (user_id, product_id, rating, comment, created_at)
VALUES
(1, 1, 5, 'Excellent smartphone, worth the price!', NOW()),
(2, 2, 4, 'Very decent laptop for work and play.', NOW()),
(3, 3, 3, 'Good sound quality but a bit heavy.', NOW()),
(4, 4, 5, 'A thrilling read, couldn’t put it down!', NOW()),
(5, 5, 4, 'Insightful book but a bit long.', NOW()),
(6, 6, 5, 'Very comfortable and well fitting!', NOW()),
(7, 7, 4, 'Stylish jeans but runs a bit small.', NOW()),
(8, 8, 5, 'Super comfy sofa, looks amazing!', NOW()),
(9, 9, 4, 'Nice design but arrived with a scratch.', NOW()),
(10, 10, 5, 'Perfect for my collection, love it!', NOW()),
(11, 11, 2, 'Deflated quickly, not as durable as expected.', NOW()),
(12, 12, 4, 'Great kit, but missing a few items.', NOW()),
(13, 13, 5, 'Best apples I\'ve ever tasted!', NOW()),
(14, 14, 4, 'Good oil for my car, will buy again!', NOW()),
(15, 15, 5, 'Incredibly useful set, great quality!', NOW()),
(16, 16, 3, 'It works fine, but nothing special.', NOW()),
(17, 17, 5, 'Stunning ring, eye-catching and beautiful!', NOW()),
(18, 18, 4, 'A refreshing drink that I truly enjoy!', NOW()),
(19, 19, 5, 'Perfect mat for my yoga sessions!', NOW()),
(20, 20, 4, 'Has all the essential tools for my repairs.', NOW());

INSERT INTO
    cart_items (user_id, item_id, quantity)
VALUES (1, 1, 1),
    (1, 2, 1),
    (2, 3, 2),
    (2, 4, 1),
    (3, 5, 1),
    (3, 6, 1),
    (4, 7, 2),
    (4, 8, 1),
    (5, 9, 3),
    (5, 10, 1),
    (6, 11, 2),
    (6, 12, 1),
    (7, 13, 5),
    (7, 14, 1),
    (8, 15, 3),
    (8, 16, 1),
    (9, 17, 1),
    (9, 18, 2),
    (10, 19, 1),
    (10, 20, 1);

INSERT INTO
    payment_methods (name)
VALUES ('Credit Card'),
    ('Debit Card'),
    ('PayPal'),
    ('Cash on Delivery'),
    ('Stripe');

INSERT INTO
    user_orders (
        user_id,
        total_amount,
        address_id,
        payment_method_id,
        payment_status,
        order_status,
        created_at,
        updated_at
    )
VALUES (
        1,
        749.98,
        1,
        1,
        'completed',
        'pending',
        NOW(),
        NOW()
    ),
    (
        2,
        1399.98,
        2,
        2,
        'pending',
        'pending',
        NOW(),
        NOW()
    ),
    (
        3,
        30.99,
        3,
        3,
        'failed',
        'cancelled',
        NOW(),
        NOW()
    ),
    (
        4,
        99.99,
        4,
        5,
        'completed',
        'shipped',
        NOW(),
        NOW()
    ),
    (
        5,
        20.99,
        5,
        6,
        'pending',
        'pending',
        NOW(),
        NOW()
    ),
    (
        6,
        179.98,
        6,
        7,
        'completed',
        'delivered',
        NOW(),
        NOW()
    ),
    (
        7,
        35.99,
        7,
        1,
        'completed',
        'pending',
        NOW(),
        NOW()
    ),
    (
        8,
        59.99,
        8,
        2,
        'pending',
        'cancelled',
        NOW(),
        NOW()
    ),
    (
        9,
        299.99,
        9,
        3,
        'completed',
        'shipped',
        NOW(),
        NOW()
    ),
    (
        10,
        479.98,
        10,
        4,
        'failed',
        'pending',
        NOW(),
        NOW()
    ),
    (
        11,
        3.99,
        11,
        5,
        'completed',
        'shipped',
        NOW(),
        NOW()
    ),
    (
        12,
        29.99,
        12,
        6,
        'pending',
        'pending',
        NOW(),
        NOW()
    ),
    (
        13,
        24.99,
        13,
        1,
        'completed',
        'delivered',
        NOW(),
        NOW()
    ),
    (
        14,
        39.99,
        14,
        2,
        'completed',
        'pending',
        NOW(),
        NOW()
    ),
    (
        15,
        89.99,
        15,
        3,
        'completed',
        'shipped',
        NOW(),
        NOW()
    ),
    (
        16,
        479.99,
        16,
        4,
        'completed',
        'delivered',
        NOW(),
        NOW()
    ),
    (
        17,
        15.99,
        17,
        5,
        'pending',
        'pending',
        NOW(),
        NOW()
    ),
    (
        18,
        9.99,
        18,
        1,
        'failed',
        'cancelled',
        NOW(),
        NOW()
    ),
    (
        19,
        25.99,
        19,
        2,
        'completed',
        'delivered',
        NOW(),
        NOW()
    ),
    (
        20,
        80.99,
        20,
        3,
        'completed',
        'shipped',
        NOW(),
        NOW()
    );