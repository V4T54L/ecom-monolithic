import React, { useState } from 'react';
import { User, Lock, Mail, UserCheck, AtSign, Loader } from 'lucide-react';
import { title } from '../constants';
import { UserInfo } from '../types';
import { Login, Signup } from '../api/user';
import { useNavigate } from 'react-router-dom';

interface AuthPageProps {
    setUserDetails: (user: UserInfo) => void
}

const AuthPage: React.FC<AuthPageProps> = ({ setUserDetails }) => {
    const [error, setError] = useState("");
    const [message, setMessage] = useState("");
    const [isLogin, setIsLogin] = useState(true);
    const [isLoading, setIsLoading] = useState(false);
    const [formData, setFormData] = useState({
        name: '',
        username: '',
        email: '',
        password: '',
    });
    const [confirmPassword, setConfirmPassword] = useState('');

    const navigate = useNavigate()

    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const { name, value } = e.target;
        setFormData({ ...formData, [name]: value });
    };

    const handleSignupSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        setError("")
        setMessage("")
        
        if (formData.password !== confirmPassword) {
            console.error("Passwords do not match");
            setError("Passwords do not match")
            return;
        }
        
        setIsLoading(true)
        try {
            const { message } = await Signup(formData.name, formData.username, formData.email, formData.password)
            setMessage(message)
            setIsLogin(true)
        } catch (err) {
            console.error("Signup error: ", err);
            setError("Signup error: " + err);
        }
        setIsLoading(false)
    };
    
    const handleLoginSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        setError("")
        setMessage("")
        
        setIsLoading(true)
        try {
            const { user } = await Login(formData.username, formData.password)
            setUserDetails(user)
            navigate("/")
        } catch (err) {
            console.error("Login error: ", err);
            setError("Login error: " + err);
        }
        setIsLoading(false)
    };



    return (
        <div className="min-h-screen bg-base-200 flex items-center justify-center py-12 px-4 sm:px-6 lg:px-8">
            <div className="max-w-md w-full space-y-8">
                <div>
                    <h2 className="text-xl border-2 w-fit mx-auto py-1 px-2 font-black rounded-3xl">{title}</h2>
                    <h2 className="mt-6 text-center text-3xl font-bold text-base-content">
                        {isLogin ? 'Log in to your account' : 'Create an account'}
                    </h2>
                    {error && (
                        <p className="text-red-500 text-center mt-2">{error}</p>
                    )}
                    {message && (
                        <p className="text-green-500 text-center mt-2">{message}</p>
                    )}
                </div>

                <form className="mt-8 space-y-6" onSubmit={isLogin ? handleLoginSubmit : handleSignupSubmit}>
                    <input type="hidden" name="remember" value="true" />

                    {!isLogin && (
                        <>
                            <div>
                                <div className="relative rounded-md shadow-sm">
                                    <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                                        <User className="h-5 w-5 text-gray-400" aria-hidden="true" />
                                    </div>
                                    <input
                                        id="name"
                                        name="name"
                                        type="text"
                                        autoComplete="name"
                                        required
                                        className="focus:ring-primary focus:border-primary block w-full pl-10 pr-3 py-2 border border-gray-300 rounded-md text-base-content placeholder-gray-500"
                                        placeholder="Full Name"
                                        value={formData.name}
                                        onChange={handleChange}
                                    />
                                </div>
                            </div>

                            <div>
                                <div className="relative rounded-md shadow-sm">
                                    <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                                        <Mail className="h-5 w-5 text-gray-400" aria-hidden="true" />
                                    </div>
                                    <input
                                        id="email"
                                        name="email"
                                        type="email"
                                        autoComplete="email"
                                        required
                                        className="focus:ring-primary focus:border-primary block w-full pl-10 pr-3 py-2 border border-gray-300 rounded-md text-base-content placeholder-gray-500"
                                        placeholder="Email address"
                                        value={formData.email}
                                        onChange={handleChange}
                                    />
                                </div>
                            </div>
                        </>
                    )}

                    <div>
                        <div className="relative rounded-md shadow-sm">
                            <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                                <AtSign className="h-5 w-5 text-gray-400" aria-hidden="true" />
                            </div>
                            <input
                                id="username"
                                name="username"
                                type="username"
                                autoComplete="current-username"
                                required
                                className="focus:ring-primary focus:border-primary block w-full pl-10 pr-3 py-2 border border-gray-300 rounded-md text-base-content placeholder-gray-500"
                                placeholder="username"
                                value={formData.username}
                                onChange={handleChange}
                            />
                        </div>
                    </div>

                    <div>
                        <div className="relative rounded-md shadow-sm">
                            <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                                <Lock className="h-5 w-5 text-gray-400" aria-hidden="true" />
                            </div>
                            <input
                                id="password"
                                name="password"
                                type="password"
                                autoComplete="current-password"
                                required
                                className="focus:ring-primary focus:border-primary block w-full pl-10 pr-3 py-2 border border-gray-300 rounded-md text-base-content placeholder-gray-500"
                                placeholder="Password"
                                value={formData.password}
                                onChange={handleChange}
                            />
                        </div>
                    </div>

                    {!isLogin && (
                        <div>
                            <div className="relative rounded-md shadow-sm">
                                <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                                    <Lock className="h-5 w-5 text-gray-400" aria-hidden="true" />
                                </div>
                                <input
                                    id="confirm-password"
                                    name="confirm-password"
                                    type="password"
                                    autoComplete="new-password"
                                    required
                                    className="focus:ring-primary focus:border-primary block w-full pl-10 pr-3 py-2 border border-gray-300 rounded-md text-base-content placeholder-gray-500"
                                    placeholder="Confirm Password"
                                    value={confirmPassword}
                                    onChange={(e) => setConfirmPassword(e.target.value)}
                                />
                            </div>
                        </div>
                    )}

                    <div>
                        <button
                            type="submit"
                            disabled={isLoading}
                            className="group relative w-full flex justify-center py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white bg-primary hover:bg-primary-focus focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary"
                        >
                            <span className="absolute left-0 inset-y-0 flex items-center pl-3">
                                {isLoading ? (
                                    <Loader className="h-5 w-5 text-white" aria-hidden="true" />
                                ) : (
                                    <UserCheck className="h-5 w-5 text-white group-hover:text-primary-content" aria-hidden="true" />
                                )}
                            </span>
                            {isLogin ? 'Sign in' : 'Sign up'}
                        </button>
                    </div>
                </form>

                <div className="mt-4 text-center text-sm text-base-content">
                    {isLogin ? (
                        <>
                            Don't have an account?{' '}
                            <button
                                onClick={() => setIsLogin(false)}
                                className="font-medium text-primary hover:text-primary-focus"
                            >
                                Sign up
                            </button>
                        </>
                    ) : (
                        <>
                            Already have an account?{' '}
                            <button
                                onClick={() => setIsLogin(true)}
                                className="font-medium text-primary hover:text-primary-focus"
                            >
                                Log in
                            </button>
                        </>
                    )}
                </div>
            </div>
        </div>
    );
};

export default AuthPage;