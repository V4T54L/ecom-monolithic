import { BrowserRouter, Route, Routes } from "react-router-dom"
import Layout from "./components/Layout"
import Home from "./components/Home"
import ViewProduct from "./components/ViewProduct"
import ViewCart from "./components/ViewCart"
import StoreContextProvider from "./contexts/StoreContext"
import Checkout from "./components/Checkout"
import ProductListingPage from "./components/ProductListingPage"

function App() {
  return (
    <>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<StoreContextProvider><Layout /></StoreContextProvider>}>
            <Route path="" element={<Home />} />
            <Route path=":id" element={<ViewProduct />} />
            <Route path="search" element={<ProductListingPage/>} />
            <Route path="cart" element={<ViewCart />} />
            <Route path="checkout" element={<Checkout />} />
            <Route path="*" element={<div>Page Not found</div>} />
          </Route>
        </Routes>
      </BrowserRouter>
    </>
  )
}

export default App
