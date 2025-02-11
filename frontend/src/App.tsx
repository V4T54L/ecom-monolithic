import { BrowserRouter, Route, Routes } from "react-router-dom"
import Layout from "./components/Layout"
import Home from "./components/Home"
import ViewProduct from "./components/ViewProduct"
import ListProduct from "./components/ListProduct"

function App() {
  return (
    <>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<Layout />}>
            <Route path="" element={<Home />} />
            <Route path="search" element={<ListProduct />} />
            <Route path=":id" element={<ViewProduct />} />
            <Route path="*" element={<div>Page Not found</div>} />
          </Route>
        </Routes>
      </BrowserRouter>
    </>
  )
}

export default App
