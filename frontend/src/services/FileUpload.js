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

  const topItems = (startDate, endDate) => {
      return http.post("/topfive", {startDate, endDate})
  }

  export default {
    upload,
    getRecords,
  };
  