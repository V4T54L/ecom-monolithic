import { BrowserRouter, Navigate, Route, Routes } from "react-router-dom"
import Layout from "./components/Layout"
import Home from "./components/Home"
import ViewProduct from "./components/ViewProduct"
import ViewCart from "./components/ViewCart"
import StoreContextProvider from "./contexts/StoreContext"
import Checkout from "./components/Checkout"
import ProductListingPage from "./components/ProductListingPage"
import UserProfilePage from "./components/UserProfilePage"
import AuthPage from "./components/AuthPage"
import { useState } from "react"
import { UserInfo } from "./types"
import instance from "./api/axios"

function App() {
  const [currentUser, setCurrentUser] = useState<UserInfo>()
  const ProtectedRoute = ({ currentUser, children }: { currentUser: UserInfo | undefined, children: React.ReactNode }) => (
    currentUser ? children : <Navigate to={"/auth"} />
  )

  instance.logout = ()=>setCurrentUser(undefined)

  return (
    <>
      <BrowserRouter>
        <Routes>
          <Route path="/auth" element={<AuthPage setUserDetails={setCurrentUser}/>} />
          <Route path="/" element={<ProtectedRoute currentUser={currentUser}><StoreContextProvider><Layout /></StoreContextProvider></ProtectedRoute>}>
            <Route path="" element={<Home />} />
            <Route path=":id" element={<ViewProduct />} />
            <Route path="profile" element={<UserProfilePage />} />
            <Route path="search" element={<ProductListingPage />} />
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
