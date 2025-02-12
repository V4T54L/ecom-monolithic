import React from 'react'

interface RatingProps {
    stars: number;
}

const Rating: React.FC<RatingProps> = ({ stars }) => {
    return (
        <>
            <div className="flex items-center my-2">
                {Array.from({ length: 5 }, (_, index) => (
                    <svg
                        key={index}
                        className={`w-5 h-5 ${index < Math.floor(stars) ? 'text-yellow-500' : 'text-gray-300'}`}
                        fill="currentColor"
                        viewBox="0 0 20 20"
                    >
                        <path d="M10 15l-5.878 3.09 1.119-6.517L0 6.18l6.545-.952L10 0l2.455 5.228 6.545.952-4.241 3.393 1.119 6.517z" />
                    </svg>
                ))}
                <span className="text-sm ml-2">{stars.toFixed(1)}</span>
            </div>
        </>
    )
}

export default Rating