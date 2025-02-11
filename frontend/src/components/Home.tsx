import React from 'react'
import Promotions from './Promotions'
import ProductCard from './ProductCard'
import { useStoreContext } from '../contexts/StoreContext'

const Home: React.FC = () => {
    const { products, loading } = useStoreContext()
    if (loading) {
        return <div>Loading...</div>
    }
    return (
        <>
            <Promotions />

            <h3 className='text-3xl font-bold'>Featured</h3>
            <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 2xl:grid-cols-4 my-4 gap-8">
                {
                    products.map(e => (
                        <ProductCard key={e.id} product={e} />
                    ))
                }
            </div>
        </>
    )
}

export default Home