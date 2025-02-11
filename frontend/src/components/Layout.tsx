import React from 'react'
import Navbar from './Navbar'
import { Outlet } from 'react-router-dom'
import Footer from './Footer'
import StoreContextProvider from '../contexts/StoreContext'

const Layout: React.FC = () => {
    return (
        <div data-theme="dark">
            <StoreContextProvider>
                <Navbar />
                <main className='min-h-screen max-w-[80%] mx-auto '>
                    <Outlet />
                </main>
                <Footer />
            </StoreContextProvider>
        </div>
    )
}

export default Layout