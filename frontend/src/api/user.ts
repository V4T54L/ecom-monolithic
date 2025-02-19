import { LoginResponse, MessageResponse, UserDetailResponse } from '../types';
import instance from './axios';

const AUTH_BASE_URI = '/auth';

const Signup = async (name: string, username: string, email: string, password: string): Promise<MessageResponse> => {
    const endpoint = `${AUTH_BASE_URI}/signup`;
    try {
        const response = await instance.axios.post<MessageResponse>(endpoint, { name, username, email, password });
        return response.data;
    } catch (error) {
        console.error("Error signing up:", error);
        throw error;
    }
};

const Login = async (username: string, password: string): Promise<LoginResponse> => {
    const endpoint = `${AUTH_BASE_URI}/login`;
    try {
        const response = await instance.axios.post<LoginResponse>(endpoint, { username, password });
        const data = response.data;
        instance.token = data.token;
        return data;
    } catch (error) {
        console.error("Error logging in:", error);
        throw error;
    }
};

const GetProfile = async (): Promise<UserDetailResponse> => {
    const endpoint = `/profile`;
    try {
        const response = await instance.axios.get<UserDetailResponse>(endpoint);
        return response.data;
    } catch (error) {
        console.error("Error logging in:", error);
        throw error;
    }
};

export { Signup, Login, GetProfile };