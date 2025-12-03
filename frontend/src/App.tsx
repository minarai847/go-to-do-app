import React, { useEffect } from 'react';
import { BrowserRouter, Routes, Route } from 'react-router-dom'
import { Auth } from './components/Auth'
import { Todo } from './components/Todo'
import axios from 'axios'
import { CsrfToken } from './types'

function App() {
  useEffect(() => {
    axios.defaults.withCredentials = true
    const getCsrfToken = async () => {
      try {
        const { data } = await axios.get<CsrfToken>(
          `${process.env.REACT_APP_API_URL}/csrf`
        )
        axios.defaults.headers.common['X-CSRF-Token'] = data.csrf_token
      } catch (error) {
        console.error('Failed to get CSRF token:', error)
      }
    }
    getCsrfToken();
  }, [])
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Auth />} />
        <Route path="/todo" element={<Todo />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
