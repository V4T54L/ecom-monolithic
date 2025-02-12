const Checkout = () => {
    return (
        <div className="container mx-auto p-6">
            <h1 className="text-3xl font-bold mb-6">Checkout</h1>

            <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                {/* Billing Information Section */}
                <div className="card bg-base-100 shadow-xl">
                    <div className="card-body">
                        <h2 className="card-title">Billing Information</h2>
                        <form>
                            <div className="form-control mb-4">
                                <label htmlFor="name" className="label">Name</label>
                                <input type="text" id="name" className="input input-bordered" placeholder="Full Name" required />
                            </div>
                            <div className="form-control mb-4">
                                <label htmlFor="email" className="label">Email</label>
                                <input type="email" id="email" className="input input-bordered" placeholder="Email Address" required />
                            </div>
                            <div className="form-control mb-4">
                                <label htmlFor="address" className="label">Shipping Address</label>
                                <input type="text" id="address" className="input input-bordered" placeholder="Address" required />
                            </div>
                            <div className="form-control mb-4">
                                <label htmlFor="city" className="label">City</label>
                                <input type="text" id="city" className="input input-bordered" placeholder="City" required />
                            </div>
                            <div className="form-control mb-4">
                                <label htmlFor="zip" className="label">Zip Code</label>
                                <input type="text" id="zip" className="input input-bordered" placeholder="Zip Code" required />
                            </div>
                        </form>
                    </div>
                </div>

                {/* Order Summary Section */}
                <div className="card bg-base-100 shadow-xl">
                    <div className="card-body">
                        <h2 className="card-title">Order Summary</h2>
                        <div className="space-y-4">
                            {/* Example Products, You can map through your products array */}
                            <div className="flex justify-between">
                                <span>Product 1</span>
                                <span>$29.99</span>
                            </div>
                            <div className="flex justify-between">
                                <span>Product 2</span>
                                <span>$19.99</span>
                            </div>
                            <div className="flex justify-between">
                                <span>Subtotal</span>
                                <span>$49.98</span>
                            </div>
                            <div className="flex justify-between">
                                <span>Shipping</span>
                                <span>$5.00</span>
                            </div>
                            <div className="flex justify-between font-bold">
                                <span>Total</span>
                                <span>$54.98</span>
                            </div>
                        </div>
                    </div>
                </div>
            </div>

            {/* Payment Section */}
            <div className="my-6">
                <h2 className="text-2xl font-bold mb-4">Payment Method</h2>
                <div className="form-control mb-4">
                    <select className="select select-bordered w-full">
                        <option disabled selected>Select Payment Method</option>
                        <option>Credit Card</option>
                        <option>PayPal</option>
                        <option>Bank Transfer</option>
                    </select>
                </div>
                <button className="btn btn-primary w-full">Complete Order</button>
            </div>
        </div>
    );
};

export default Checkout;