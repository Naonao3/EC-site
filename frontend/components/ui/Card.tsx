import React from 'react'

interface CardProps {
  children: React.ReactNode
  className?: string
  hover?: boolean
  onClick?: () => void
}

export const Card: React.FC<CardProps> = ({
  children,
  className = '',
  hover = false,
  onClick,
}) => {
  const hoverClass = hover ? 'hover:shadow-lg hover:scale-105 transition-all duration-200 cursor-pointer' : ''

  return (
    <div
      className={`bg-white rounded-lg shadow-md p-6 ${hoverClass} ${className}`}
      onClick={onClick}
    >
      {children}
    </div>
  )
}
