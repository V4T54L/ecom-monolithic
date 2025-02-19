interface ProductInfo {
    id: number;
    thumbnail: string;
    title: string;
    description: string;
    rating: number;
    price: number;
    discountedPrice?: number;
}

interface Review {
    reviewerName: string;
    rating: number;
    comment: string;
    date: string;
}

interface ProductDetail extends ProductInfo {
    images: string[];
    category: string;
    stock: number;
    specifications: {
        [key: string]: string | number;
    };
    reviews: Review[];
    relatedProducts?: ProductInfo[];
}

interface CartItem extends ProductInfo {
    quantity: number
}

interface Cart {
    totalAmount: number;
    items: CartItem[];
}

interface Address {
    address: string;
    city: string;
    zipCode: string;
}

interface User {
    id?: number;
    name: string;
    username: string;
    email: string;
    role: string;
    memberSince: string;
    addresses: Address[];
}

interface UserInfo {
    id: number;
    name: string;
    username: string;
    email: string;
    role: string;
    created_at: string;
    updated_at: string;
}

interface UserDetailResponse {
    user: UserInfo;
}

interface LoginResponse {
    user: UserInfo;
    token: string;
}

interface MessageResponse {
    message: string;
}

interface ErrorResponse {
    error: string;
}


export type {
    ProductInfo, Review, ProductDetail, Cart,
    CartItem, User, Address, UserDetailResponse, ErrorResponse,
    UserInfo, MessageResponse, LoginResponse,
}