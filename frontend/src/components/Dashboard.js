import React from 'react'

function Dashboard({ records }) {
    return (
        <div>
            <div className="date-picker"></div>

            {/* dashboards */}
            <div className="dashboards">
                <div className="records">
                    <table className="table table-striped">
                        <thead>
                            <tr>
                                <th scope="col">OrderID</th>
                                <th scope="col">ItemType</th>
                                <th scope="col">OrderDate</th>
                            <th scope="col">TotalProfit</th>
                            </tr>
                        </thead>
                        <tbody>
                            {records.map(record =>
                                <tr>
                                    <td>{record.OrderID}</td>
                                    <td>{record.ItemType}</td>
                                    <td>{record.OrderDate}</td>
                                    <td>{record.TotalProfit}</td>
                                </tr>
                            )}
                        </tbody>
                    </table>
                </div>
                
                <div className="profits">Hello</div>
            </div>
        </div>
    )
}

export default Dashboard
