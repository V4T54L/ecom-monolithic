import { ProductDetail, ProductInfo } from "../types";

const mockProducts: ProductInfo[] = [
    {
        id: 1,
        thumbnail: 'https://placehold.co/600x400?text=Wired+Headphones',
        title: 'Wired Headphones',
        description: 'Comfortable over-ear wired headphones with exceptional sound quality.',
        rating: 3.5,
        price: 39.99,
        discountedPrice: 29.99,
    },
    {
        id: 2,
        thumbnail: 'https://placehold.co/600x400?text=Laptop',
        title: 'Gaming Laptop',
        description: 'High-performance gaming laptop with stunning graphics and speed.',
        rating: 4.8,
        price: 1499.99,
    },
    {
        id: 3,
        thumbnail: 'https://placehold.co/600x400?text=Smartphone',
        title: 'Smartphone ABC',
        description: 'Latest smartphone with all advanced features and a stunning display.',
        rating: 4.6,
        price: 799.99,
        discountedPrice: 749.99,
    },
    {
        id: 4,
        thumbnail: 'https://placehold.co/600x400?text=Bluetooth+Speaker',
        title: 'Portable Bluetooth Speaker',
        description: 'Waterproof Bluetooth speaker with deep bass for outdoor fun.',
        rating: 4.7,
        price: 59.99,
        discountedPrice: 49.99,
    },
    {
        id: 5,
        thumbnail: 'https://placehold.co/600x400?text=Smartwatch',
        title: 'Smartwatch Z',
        description: 'Feature-rich smartwatch with heart rate monitor and fitness tracking.',
        rating: 4.3,
        price: 199.99,
    },
    {
        id: 6,
        thumbnail: 'https://placehold.co/600x400?text=Gaming+Mouse',
        title: 'RGB Gaming Mouse',
        description: 'High precision gaming mouse with customizable RGB lighting.',
        rating: 4.2,
        price: 59.99,
        discountedPrice: 49.99,
    },
    {
        id: 7,
        thumbnail: 'https://placehold.co/600x400?text=4K+Monitor',
        title: '4K Ultra HD Monitor',
        description: 'Stunning 4K monitor with a wide color gamut and high refresh rate.',
        rating: 4.9,
        price: 499.99,
    },
    {
        id: 8,
        thumbnail: 'https://placehold.co/600x400?text=Wireless+Charger',
        title: 'Wireless Charger',
        description: 'Fast wireless charger compatible with all Qi-enabled devices.',
        rating: 4.4,
        price: 29.99,
    },
];


const mockProductDetail: ProductDetail = {
    id: 1,
    title: 'Wireless Headphones',
    description: 'High quality wireless headphones with noise cancellation.',
    rating: 4.7,
    price: 89.99,
    discountedPrice: 79.99,
    thumbnail: 'https://example.com/images/headphones1.jpg',
    images: [
        'https://placehold.co/600x600?text=Wired+Headphones+One',
        'https://placehold.co/600x600?text=Wired+Headphones+TWO',
        'https://placehold.co/600x600?text=Wired+Headphones+THREE',
    ],
    category: 'Electronics',
    stock: 150,
    specifications: {
        batteryLife: '20 hours',
        weight: '0.3 kg',
        color: 'Black',
        connectivity: 'Bluetooth 5.0',
    },
    reviews: [
        {
            reviewerName: 'Alice Johnson',
            rating: 5,
            comment: 'Absolutely fantastic sound quality and comfort!',
            date: '2023-10-01',
        },
        {
            reviewerName: 'Bob Smith',
            rating: 4,
            comment: 'Great headphones, but some issues with Bluetooth range.',
            date: '2023-10-05',
        },
    ],
    relatedProducts: [
        {...mockProducts[0]},
        {...mockProducts[1]},
        {...mockProducts[2]},
    ],
};

export { mockProductDetail, mockProducts };