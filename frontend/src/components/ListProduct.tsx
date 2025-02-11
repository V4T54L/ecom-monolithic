import React from 'react'
import ProductCard from './ProductCard'
import { useStoreContext } from '../contexts/StoreContext'

const ListProduct: React.FC = () => {
    const { products } = useStoreContext()
    return (
        <>
            <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 2xl:grid-cols-4 my-4 gap-8">
                {
                    products.map(e => (
                        <ProductCard key={e.id} product={e} />
                    ))
                }
            </div>
            <div className="join grid grid-cols-2">
                <button className="join-item btn btn-outline">Previous page</button>
                <button className="join-item btn btn-outline">Next</button>
            </div>
        </>
    )
}

export default ListProduct