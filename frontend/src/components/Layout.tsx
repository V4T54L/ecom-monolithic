import React from 'react'
import Navbar from './Navbar'
import { Outlet } from 'react-router-dom'
import Footer from './Footer'
import { useStoreContext } from '../contexts/StoreContext'

const Layout: React.FC = () => {
    const { isDarkMode } = useStoreContext();
    return (
        <div data-theme={isDarkMode ? "dark" : "light"} className=''>
            <Navbar />
            <main className='container px-[5%] mx-auto min-h-screen'>
                <Outlet />
            </main>
            <Footer />
        </div >
    )
}

export default Layout