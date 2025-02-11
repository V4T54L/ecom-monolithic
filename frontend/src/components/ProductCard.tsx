import React from 'react'
import { ProductInfo } from '../types'
import { useStoreContext } from '../contexts/StoreContext'
import { useNavigate } from 'react-router-dom'

interface ProductCardProps {
    product: ProductInfo
}

const ProductCard: React.FC<ProductCardProps> = ({ product }) => {
    const { addProductToCart } = useStoreContext()

    const navigate = useNavigate()

    const viewProduct = (id: number) => {
        navigate(`/${id}`)
    }

    return (
        <div className="card card-compact bg-base-100 min-w-60 shadow-xl">
            <figure>
                <img
                    src={product.thumbnail}
                    alt={product.title} />
            </figure>
            <div className="card-body">
                <h2 className="card-title text-lg">{product.title}</h2>
                <p className='text-xs'>{product.description}</p>

                <div className="flex items-center my-2">
                    {/* Generate filled stars based on the rating */}
                    {Array.from({ length: 5 }, (_, index) => (
                        <svg
                            key={index}
                            className={`w-5 h-5 ${index < Math.floor(product.rating) ? 'text-yellow-500' : 'text-gray-300'}`}
                            fill="currentColor"
                            viewBox="0 0 20 20"
                        >
                            <path d="M10 15l-5.878 3.09 1.119-6.517L0 6.18l6.545-.952L10 0l2.455 5.228 6.545.952-4.241 3.393 1.119 6.517z" />
                        </svg>
                    ))}
                    {/* Display the rating number */}
                    <span className="text-sm ml-2">{product.rating.toFixed(1)}</span>
                </div>
                {
                    product.discountedPrice ? (
                        <div className='flex items-end gap-2'>
                            <div className="text-sm line-through font-extralight">${product.price}</div>
                            <div className="text-xl">${product.discountedPrice}</div>
                        </div>
                    ) : (
                        <div className="text-xl">${product.price}</div>
                    )
                }

                <div className="card-actions justify-between">
                    <button onClick={() => viewProduct(product.id)} className="btn btn-ghost">View Product</button>
                    <button onClick={() => addProductToCart(product)} className="btn btn-primary">Add to cart</button>
                </div>
            </div>
        </div>
    )
}

export default ProductCard
