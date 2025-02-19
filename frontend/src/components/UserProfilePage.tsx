// src/components/UserProfilePage.tsx

import React, { useState, useEffect } from 'react';
import { Address, User } from '../types';
import { Pencil, Save, XCircle, Plus, Trash2 } from 'lucide-react';
import { mockUser } from '../mock/user';
import { GetProfile } from '../api/user';


const UserProfilePage: React.FC = () => {
    const [user, setUser] = useState<User | null>(null);
    const [loading, setLoading] = useState<boolean>(true);
    const [error, setError] = useState<string | null>(null);
    const [editMode, setEditMode] = useState<boolean>(false);
    const [tempUser, setTempUser] = useState<User | null>(null);

    useEffect(() => {
        const fetchUserProfile = async () => {
            try {
                setLoading(true);
                const { user } = await GetProfile()
                setUser({
                    id: user.id, email: user.email, name: user.name,
                    memberSince: user.created_at, role: user.role,
                    username: user.username, addresses: [],
                });
                setTempUser(mockUser);
                setError(null);
            } catch (err) {
                if (err instanceof Error) {
                    setError(err.message);
                } else {
                    setError('Failed to fetch user profile.');
                }
            } finally {
                setLoading(false);
            }
        };

        fetchUserProfile();
    }, []);

    const handleInputChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
        if (!tempUser) return;
        const { name, value } = e.target;
        setTempUser({ ...tempUser, [name]: value });
    };

    const handleAddressChange = (index: number, field: keyof Address, value: string) => {
        if (!tempUser) return;

        const updatedAddresses = [...tempUser.addresses];
        updatedAddresses[index] = { ...updatedAddresses[index], [field]: value };

        setTempUser({ ...tempUser, addresses: updatedAddresses });
    };

    const handleAddAddress = () => {
        if (!tempUser) return;
        setTempUser({
            ...tempUser,
            addresses: [...tempUser.addresses, { address: '', city: '', zipCode: '' }],
        });
    };

    const handleRemoveAddress = (index: number) => {
        if (!tempUser) return;
        const updatedAddresses = [...tempUser.addresses];
        updatedAddresses.splice(index, 1);
        setTempUser({ ...tempUser, addresses: updatedAddresses });
    };

    const handleSave = async () => {
        if (!tempUser) return;
        try {
            setLoading(true);

            // TODO: make API call

            setUser(tempUser);
            setEditMode(false);
            setError(null);
        } catch (err) {
            if (err instanceof Error) {
                setError(err.message || 'Failed to update user profile.');
            } else {
                setError('Failed to update user profile.');
            }
        } finally {
            setLoading(false);
        }
    };

    const handleCancel = () => {
        setEditMode(false);
        setTempUser(user ? { ...user } : null);
    };

    if (loading) {
        return <div className="text-center py-4">Loading user profile...</div>;
    }

    if (error) {
        return <div className="text-red-500 text-center py-4">Error: {error}</div>;
    }

    if (!user) {
        return <div className="text-center py-4">User profile not found.</div>;
    }

    return (
        <div className="container mx-auto p-4">
            <div className="flex items-center justify-between mb-4">
                <h1 className="text-2xl font-bold">User Profile</h1>
                {editMode ? (
                    <div>
                        <button className="btn btn-sm btn-success mr-2" onClick={handleSave} disabled={loading}>
                            <Save className="mr-1" size={16} />
                            Save
                        </button>
                        <button className="btn btn-sm btn-ghost" onClick={handleCancel} disabled={loading}>
                            <XCircle className="mr-1" size={16} />
                            Cancel
                        </button>
                    </div>
                ) : (
                    <button className="btn btn-sm btn-primary" onClick={() => setEditMode(true)}>
                        <Pencil className="mr-1" size={16} />
                        Edit Profile
                    </button>
                )}
            </div>

            <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                {/* Basic Information */}
                <div className="card bg-base-100 shadow-md">
                    <div className="card-body">
                        <h2 className="card-title">Basic Information</h2>
                        <div className="form-control">
                            <label className="label">
                                <span className="label-text">Name:</span>
                            </label>
                            <input
                                type="text"
                                name="name"
                                value={tempUser?.name || ''}
                                onChange={handleInputChange}
                                className="input input-bordered"
                                disabled={!editMode}
                            />
                        </div>
                        <div className="form-control">
                            <label className="label">
                                <span className="label-text">Username:</span>
                            </label>
                            <input
                                type="text"
                                name="username"
                                value={user.username}
                                className="input input-bordered"
                                disabled
                            />
                        </div>
                        <div className="form-control">
                            <label className="label">
                                <span className="label-text">Email:</span>
                            </label>
                            <input
                                type="email"
                                name="email"
                                value={tempUser?.email || ''}
                                onChange={handleInputChange}
                                className="input input-bordered"
                                disabled={!editMode}
                            />
                        </div>
                    </div>
                </div>

                {/* Address Information */}
                <div className="card bg-base-100 shadow-md">
                    <div className="card-body">
                        <h2 className="card-title">Addresses</h2>
                        {tempUser?.addresses.map((address, index) => (
                            <div key={index} className="mb-4 border p-2 rounded">
                                <h3 className="text-lg font-semibold">Address #{index + 1}</h3>
                                <div className="form-control">
                                    <label className="label">
                                        <span className="label-text">Address:</span>
                                    </label>
                                    <input
                                        type="text"
                                        name="address"
                                        value={address.address}
                                        onChange={(e) => handleAddressChange(index, 'address', e.target.value)}
                                        className="input input-bordered"
                                        disabled={!editMode}
                                    />
                                </div>
                                <div className="form-control">
                                    <label className="label">
                                        <span className="label-text">City:</span>
                                    </label>
                                    <input
                                        type="text"
                                        name="city"
                                        value={address.city}
                                        onChange={(e) => handleAddressChange(index, 'city', e.target.value)}
                                        className="input input-bordered"
                                        disabled={!editMode}
                                    />
                                </div>
                                <div className="form-control">
                                    <label className="label">
                                        <span className="label-text">Zip Code:</span>
                                    </label>
                                    <input
                                        type="text"
                                        name="zipCode"
                                        value={address.zipCode}
                                        onChange={(e) => handleAddressChange(index, 'zipCode', e.target.value)}
                                        className="input input-bordered"
                                        disabled={!editMode}
                                    />
                                </div>
                                {editMode && (
                                    <button className="btn btn-sm btn-error mt-2" onClick={() => handleRemoveAddress(index)}>
                                        <Trash2 className="mr-1" size={16} />
                                        Remove
                                    </button>
                                )}
                            </div>
                        ))}
                        {editMode && (
                            <button className="btn btn-sm btn-primary" onClick={handleAddAddress}>
                                <Plus className="mr-1" size={16} />
                                Add Address
                            </button>
                        )}
                    </div>
                </div>

                {/* Read-Only Information */}
                <div className="card bg-base-100 shadow-md">
                    <div className="card-body">
                        <h2 className="card-title">Account Information</h2>
                        <div className="form-control">
                            <label className="label">
                                <span className="label-text">Role:</span>
                            </label>
                            <input
                                type="text"
                                name="role"
                                value={user.role}
                                className="input input-bordered"
                                disabled
                            />
                        </div>
                        <div className="form-control">
                            <label className="label">
                                <span className="label-text">Member Since:</span>
                            </label>
                            <input
                                type="text"
                                name="memberSince"
                                value={user.memberSince}
                                className="input input-bordered"
                                disabled
                            />
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default UserProfilePage;