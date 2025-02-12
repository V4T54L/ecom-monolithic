import React from 'react'
import { useStoreContext } from '../contexts/StoreContext'
import { useNavigate } from 'react-router-dom'
import Rating from './Rating'

const ViewProduct: React.FC = () => {
    const { productDetail, loading, error, addProductToCart } = useStoreContext()

    const navigate = useNavigate()

    const viewProduct = (id: number) => {
        navigate(`/${id}`)
    }

    if (loading) {
        return <div>Loading...</div>
    }

    if (error) {
        return <div>Error occured: {error}</div>
    }

    if (!productDetail) {
        return <div>Product not found</div>
    }


    return (
        <>
            <div className="max-w-4xl mx-auto p-6">
                <div className="flex">
                    {/* Image carousal */}
                    <div className="carousel rounded-box w-64">
                        {
                            productDetail?.images.map(url => (
                                <div key={url} className="carousel-item w-full">
                                    <img
                                        src={url}
                                        className="w-full"
                                        alt="Tailwind CSS Carousel component" />
                                </div>
                            ))
                        }
                    </div>

                    <div className="ml-6">
                        <h1 className="text-4xl font-bold">{productDetail.title}</h1>
                        <p className="text-gray-600 text-xl mt-2">
                            {productDetail.discountedPrice ? (
                                <>
                                    <span className="line-through text-primary">${productDetail.price.toFixed(2)}</span>
                                    <span className="text-secondary text-2xl ml-2">${productDetail.discountedPrice.toFixed(2)}</span>
                                </>
                            ) : (
                                <span className="text-2xl">${productDetail.price.toFixed(2)}</span>
                            )}
                        </p>
                        <Rating stars={productDetail.rating} />

                        <p className="mt-4">{productDetail.description}</p>
                        <button onClick={() => addProductToCart(productDetail)} className="btn btn-primary mt-8">Add to cart</button>
                    </div>
                </div>
                <div className="mt-6">
                    <h2 className="text-2xl font-semibold">Specifications</h2>

                    {/* Specifications table */}
                    <div className="overflow-x-auto">
                        <table className="table table-zebra">
                            {/* head */}
                            <thead>
                                <tr>
                                    <th>Key</th>
                                    <th>Value</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    Object.entries(productDetail.specifications).map(([key, value]) => (
                                        <tr key={key}>
                                            <td>{key}</td>
                                            <td>{value}</td>
                                        </tr>
                                    ))
                                }
                            </tbody>
                        </table>
                    </div>

                </div>
                <div className="mt-6">
                    <h2 className="text-2xl font-semibold">Reviews</h2>
                    {productDetail.reviews.length === 0 ? (
                        <p>No reviews available.</p>
                    ) : (
                        productDetail.reviews.map((review, index) => (
                            <div key={index} className="border p-4 rounded-lg mt-4">
                                <p className="text-gray-500"><strong>{review.reviewerName}</strong> - {review.date}</p>
                                <Rating stars={review.rating} />
                                <p>{review.comment}</p>
                            </div>
                        ))
                    )}
                </div>
                {productDetail.relatedProducts && productDetail.relatedProducts.length > 0 && (
                    <div className="mt-6">
                        <h2 className="text-2xl font-semibold">Related Products</h2>
                        <div className="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6 mt-4">
                            {productDetail.relatedProducts.map((relatedProduct) => (
                                <div key={relatedProduct.id} className="border rounded-lg p-4 flex flex-col h-full">
                                    <img src={relatedProduct?.thumbnail} alt={relatedProduct.title} className="w-full h-40 object-cover rounded-t-lg" />
                                    <h3 className="flex-1 text-xl font-semibold mt-2">{relatedProduct.title}</h3>
                                    {
                                        relatedProduct.discountedPrice ? (
                                            <div className='flex items-end gap-2'>
                                                <div className="text-sm line-through font-extralight">${relatedProduct.price}</div>
                                                <div className="text-xl">${relatedProduct.discountedPrice}</div>
                                            </div>
                                        ) : (
                                            <div className="text-xl">${relatedProduct.price}</div>
                                        )
                                    }
                                    <Rating stars={relatedProduct.rating} />
                                    <button onClick={() => viewProduct(relatedProduct.id)} className="btn btn-primary w-full">View Product</button>
                                </div>

                            ))}
                        </div>
                    </div>
                )}
            </div>

        </>
    )
}

export default ViewProduct