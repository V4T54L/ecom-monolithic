import { useStoreContext } from "../contexts/StoreContext";

const ViewCart: React.FC = () => {
    const { cart, removeProductFromCart } = useStoreContext()
    return (
        <div className="max-w-3xl mx-auto p-4">
            <h1 className="text-2xl font-bold mb-4">Shopping Cart</h1>
            {cart.items.length === 0 ? (
                <p>Your cart is empty</p>
            ) : (
                <>
                    <ul className="rounded-lg shadow-md divide-y divide-neutral-content">
                        {cart.items.map((item) => (
                            <li key={item.id} className="flex items-center justify-between p-4">
                                <div className="flex items-center space-x-4">
                                    <img src={item.thumbnail} alt={item.title} className="w-20 h-20 object-cover rounded-md" />
                                    <div className="flex-1">
                                        <h2 className="text-lg font-semibold">{item.title}</h2>
                                        <p className="text-sm text-gray-600">{item.description}</p>
                                        <div className="flex items-center space-x-2 mt-1">
                                            <span className="text-lg font-bold">
                                                {item.discountedPrice ? (
                                                    <span className="line-through text-gray-600">{item.price.toFixed(2)} </span>
                                                ) : (
                                                    item.price.toFixed(2)
                                                )}
                                                <span className="text-red-600">{item.discountedPrice ? item.discountedPrice.toFixed(2) : item.price.toFixed(2)}</span>
                                            </span>
                                            <span className="text-sm text-gray-500">x {item.quantity}</span>
                                        </div>
                                    </div>
                                </div>
                                <button
                                    onClick={() => removeProductFromCart(item)}
                                    className="text-red-500 hover:text-red-700"
                                >
                                    Remove
                                </button>
                            </li>
                        ))}
                    </ul>
                    <div className="mt-4 flex justify-between items-center">
                        <span className="text-xl font-semibold">Total Amount: ${cart.totalAmount.toFixed(2)}</span>
                        <button
                            // onClick={onCheckout}
                            className="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700"
                        >
                            Checkout
                        </button>
                    </div>
                </>
            )}
        </div>
    );
};

export default ViewCart;
