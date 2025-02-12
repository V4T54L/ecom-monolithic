import React, { FormEvent, useEffect, useState } from 'react'
import { title } from '../constants'
import { useStoreContext } from '../contexts/StoreContext'
import { useNavigate } from 'react-router-dom'
import ThemeProvider from './ThemeProvider'

const Navbar: React.FC = () => {

    const { cart } = useStoreContext()
    const [itemCount, setItemCount] = useState(0)
    const navigate = useNavigate()
    const [query, setQuery] = useState("")

    useEffect(() => {
        console.log("Items : ", cart.items)
        const quantity = cart.items.reduce((total, item) => total + item.quantity, 0)
        setItemCount(quantity)
    }, [cart])

    const handleFormSubmit = (e: FormEvent<HTMLFormElement>) => {
        e.preventDefault()
        navigate(`/search?q=${query}`)
    }

    return (
        <div className="navbar bg-base-100 sticky top-0 z-10">
            <div className="flex-1">
                <button onClick={() => navigate("/")} className="btn btn-ghost text-xl">{title}</button>
            </div>
            <form onSubmit={handleFormSubmit} className="flex-1 form-control">
                <input type="text" placeholder="Search" value={query} onChange={(e) => setQuery(e.target.value)} className="input input-bordered w-24 md:w-full" />
            </form>
            <div className="mx-8">
                <ThemeProvider />
            </div>
            <div className="flex-none gap-2">
                <div className="dropdown dropdown-end">
                    <div tabIndex={0} role="button" className="btn btn-ghost btn-circle">
                        <div className="indicator">
                            <svg
                                xmlns="http://www.w3.org/2000/svg"
                                className="h-5 w-5"
                                fill="none"
                                viewBox="0 0 24 24"
                                stroke="currentColor">
                                <path
                                    strokeLinecap="round"
                                    strokeLinejoin="round"
                                    strokeWidth="2"
                                    d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z" />
                            </svg>
                            <span className="badge badge-sm indicator-item">{itemCount}</span>
                        </div>
                    </div>
                    <div
                        tabIndex={0}
                        className="card card-compact dropdown-content bg-base-100 z-[1] mt-3 w-52 shadow">
                        <div className="card-body">
                            <span className="text-lg font-bold">{itemCount} Items</span>
                            <span className="text-info">Subtotal: ${cart.totalAmount}</span>
                            <div className="card-actions">
                                <button onClick={() => navigate("/cart")} className="btn btn-primary btn-block">View cart</button>
                            </div>
                        </div>
                    </div>
                </div>
                <div className="dropdown dropdown-end">
                    <div tabIndex={0} role="button" className="btn btn-ghost btn-circle avatar">
                        <div className="w-10 rounded-full">
                            <img
                                alt="Tailwind CSS Navbar component"
                                src="https://img.daisyui.com/images/stock/photo-1534528741775-53994a69daeb.webp" />
                        </div>
                    </div>
                    <ul
                        tabIndex={0}
                        className="menu menu-sm dropdown-content bg-base-100 rounded-box z-[1] mt-3 w-52 p-2 shadow">
                        <li>
                            <a className="justify-between">
                                Profile
                                <span className="badge">New</span>
                            </a>
                        </li>
                        <li><a>Settings</a></li>
                        <li><a>Logout</a></li>
                    </ul>
                </div>

            </div>
        </div>
    )
}

export default Navbar