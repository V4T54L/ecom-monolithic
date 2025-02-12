import React from 'react'
import { ProductInfo } from '../types'
import { useStoreContext } from '../contexts/StoreContext'
import { useNavigate } from 'react-router-dom'
import Rating from './Rating'

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

                <Rating stars={product.rating} />
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
