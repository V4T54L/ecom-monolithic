import React, { createContext, useContext, useEffect, useState } from 'react'
import { Cart, CartItem, ProductDetail, ProductInfo } from '../types';
import { mockProductDetail, mockProducts } from '../mock/products';


interface StoreContextProps {
    products: ProductInfo[];
    productDetail: ProductDetail | undefined;
    loading: boolean; error: string | null; cart: Cart;
    addProductToCart: (product: ProductInfo) => void;
    removeProductFromCart: (product: ProductInfo) => void;
    isDarkMode: boolean;
    themeToggleFunc: () => void;
}

const StoreContext = createContext<StoreContextProps | null>(null)

interface StoreContextProviderProps {
    children: React.ReactNode
}
const StoreContextProvider: React.FC<StoreContextProviderProps> = ({ children }) => {
    const [products, setProducts] = useState<ProductInfo[]>([]);
    const [productDetail, setProductDetail] = useState<ProductDetail>();
    const [loading, setLoading] = useState<boolean>(true);
    const [isDarkMode, setDarkMode] = useState<boolean>(true);
    const [error, setError] = useState<string | null>(null);
    const [cart, setCart] = useState<Cart>({
        items: [],
        totalAmount: 0,
    })

    const updateTotalAmount = (items: CartItem[]) => {
        return Math.round(items.reduce((total, item) => total + item.quantity * (item.discountedPrice ? item.discountedPrice : item.price), 0) * 1000) / 1000;
    };

    const addProductToCart = (product: ProductInfo) => {
        console.log("Adding : ", product.title)
        setCart(prev => {
            const existingCartItemIndex = prev.items.findIndex(item => item.id === product.id);

            const updatedItems = [...prev.items];
            if (existingCartItemIndex > -1) {
                updatedItems[existingCartItemIndex].quantity += 1;
            } else {
                updatedItems.push({ ...product, quantity: 1 })
            }
            return {
                items: updatedItems,
                totalAmount: updateTotalAmount(updatedItems),
            };
        });
    };

    const themeToggleFunc = () => {
        setDarkMode(prev => !prev)
    }

    const removeProductFromCart = (product: ProductInfo) => {
        setCart(prev => {
            const existingCartItemIndex = prev.items.findIndex(item => item.id === product.id);

            if (existingCartItemIndex > -1) {
                const updatedItems = [...prev.items];
                const itemToUpdate = updatedItems[existingCartItemIndex];

                if (itemToUpdate.quantity > 1) {
                    itemToUpdate.quantity -= 1;
                    return {
                        items: updatedItems,
                        totalAmount: prev.totalAmount - (product.discountedPrice ? product.discountedPrice : product.price),
                    };
                } else {
                    updatedItems.splice(existingCartItemIndex, 1);
                    return {
                        items: updatedItems,
                        totalAmount: prev.totalAmount - (product.discountedPrice ? product.discountedPrice : product.price),
                    };
                }
            }

            return prev;
        });
    };



    useEffect(() => {
        const fetchProducts = async () => {
            try {
                setLoading(true);
                await new Promise((resolve) => setTimeout(resolve, 1000));

                setProducts(mockProducts);

                setProductDetail(mockProductDetail);
            } catch (err) {
                setError('Failed to fetch products : ' + err);
            } finally {
                setLoading(false);
            }
        };

        fetchProducts();
    }, []);

    return (
        <StoreContext.Provider value={{
            products,
            productDetail,
            loading,
            error,
            cart,
            addProductToCart,
            removeProductFromCart,
            isDarkMode,
            themeToggleFunc,
        }}>
            {children}
        </StoreContext.Provider>
    )
}

export function useStoreContext() {
    const context = useContext(StoreContext);
    if (!context) {
        throw new Error(
            "useStoreContext must be used within StoreContextProvider"
        );
    }

    return context
}

export default StoreContextProvider