import React from 'react'

function TopProfitableItems({ items }) {
    return (
        <div>
            <p>Top Five Profitable</p>
            <table className="table">
            <thead className="thead-dark">
                <tr>
                    <th scope="col">Name</th>
                    <th scope="col">Profit</th>
                </tr>
            </thead>
            <tbody>
                {items.map(item =>
                <tr>
                    <td>{item.name}</td>
                    <td>{item.profit}</td>
                </tr>
                )}
            </tbody>
            </table>
        </div>
    )
}

export default TopProfitableItems
