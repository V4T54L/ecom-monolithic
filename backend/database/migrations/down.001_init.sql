DROP TABLE IF EXISTS user_orders;

DROP TABLE IF EXISTS payment_methods;

DROP TABLE IF EXISTS cart_items;

DROP TABLE IF EXISTS product_reviews;

DROP TABLE IF EXISTS product_images;

DROP TRIGGER IF EXISTS update_products_updated_at_trigger ON products;

DROP FUNCTION IF EXISTS update_products_updated_at ();

DROP TABLE IF EXISTS products;

DROP TABLE IF EXISTS categories;

DROP TABLE IF EXISTS addresses;

DROP TRIGGER IF EXISTS update_users_updated_at_trigger ON users;

DROP FUNCTION IF EXISTS update_users_updated_at ();

DROP TABLE IF EXISTS users;

DROP TABLE IF EXISTS versions;

DROP TYPE IF EXISTS payment_status_enum;

DROP TYPE IF EXISTS order_status_enum;