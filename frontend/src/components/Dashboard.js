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
                                <th scope="col">Order Date</th>
                                <th scope="col">Order Priority</th>
                                <th scope="col">Units Sold</th>
                                <th scope="col">Unit Price</th>
                                <th scope="col">Total Cost</th>
                                <th scope="col">Total Revenue</th>
                                <th scope="col">Item Type</th>
                            </tr>
                        </thead>
                        <tbody>
                            {records.map(record =>
                                <tr>
                                    <td>{new Date(record.OrderDate).toISOString()}</td>
                                    <td>{record.OrderPriority}</td>
                                    <td>{record.UnitsSold}</td>
                                    <td>{record.UnitPrice}</td>
                                    <td>{record.TotalCost}</td>
                                    <td>{record.TotalRevenue}</td>
                                    <td>{record.ItemType}</td>
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
