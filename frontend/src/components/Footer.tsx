import React from 'react'
import { title } from '../constants'

const Footer: React.FC = () => {
    return (
        <footer className="footer footer-center bg-base-300 text-base-content p-4">
            <aside>
                <p>Copyright © {new Date().getFullYear()} - All right reserved by {title}</p>
            </aside>
        </footer>
    )
}

export default Footer