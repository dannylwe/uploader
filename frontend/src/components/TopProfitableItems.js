import React, { useState, useEffect } from 'react';
import DatePicker from "react-datepicker";
import UploadService from "../services/FileUpload"
import "react-datepicker/dist/react-datepicker.css";

function TopProfitableItems({ items }) {
    const [startDate, setStartDate] = useState(new Date());
    const [itemsTop, setItems] = useState([]);


        async function getItems (searchDate) {
            await UploadService.topItems(searchDate, searchDate).then((response) => {
                setItems(response.data);
                // console.log(response.data)
                console.log("call from component")
            });
        }
        let searchDate = new Date(startDate).toISOString().split('T')[0]
        // console.log(itemsTop)
     

    return (
        <>
            <div className="profitable">
                <p>Top Five Profitable</p>
                <DatePicker 
                selected={startDate} 
                onChange={date => {
                    setStartDate(date)
                    getItems(searchDate)
                }
                } />

                <table className="table">
                <thead className="thead-dark">
                    <tr>
                        <th scope="col">Name</th>
                        <th scope="col">Profit</th>
                    </tr>
                </thead>
                <tbody>
                    {itemsTop != null && itemsTop.map(items =>
                    <tr>
                        <td>{items.name}</td>
                        <td>{items.profit}</td>
                    </tr>
                    )}

                    {items && items.map(item =>
                    <tr>
                        <td>{item.name}</td>
                        <td>{item.profit}</td>
                    </tr>
                    )}

                    
                </tbody>
                </table>
            </div>
        </>
    )
}

export default TopProfitableItems
