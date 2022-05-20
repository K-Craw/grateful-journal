import React from "react"
const Header = () => {
    return (
        <div className = 'headerBar'>
            <div className = 'entryBar'>
                <p> Add Entry </p>
                <p> Previous Entries </p>
            </div>
            <div className = 'userBar'>
                <h2 className="username">UserNameDummy</h2>
            </div>
        </div>
    )
}

export default Header