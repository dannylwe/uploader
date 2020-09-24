import Axios from "axios";
import http from "../httpCommon";

const upload = (file, onUploadProgress) => {
    let formData = new FormData();
  
    formData.append("myFile", file);
  
    return http.post("/upload", formData, {
      headers: {
        "Content-Type": "multipart/form-data",
      },
      onUploadProgress,
    });
  };
  
  const getRecords = () => {
    return http.get("/records");
  };

  const topItems = async (startDate, endDate) => {
      let data = {startDate, endDate}
      console.log(data)
      let items = await Axios.post("http://localhost:8080/topfive", data)
      return items
  }

  const profitByDate = async (startDate, endDate) => {
    let data = {startDate, endDate}
    console.log(data)
    let items = await Axios.post("http://localhost:8080/profit", data)
    return items
}

  export default {
    upload,
    getRecords,
    topItems,
    profitByDate
  };
  