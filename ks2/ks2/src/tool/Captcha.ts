export function generateCaptcha() {
    const chars =
      "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789";
    let text = "";
    for (let i = 0; i < 4; i++) {
      text += chars.charAt(Math.floor(Math.random() * chars.length));
    }
   let captchaText = text.toLowerCase(); // 转换为小写
  
    const canvas = document.querySelector("canvas");
    if (canvas) {
      const ctx = canvas.getContext("2d");
      if (ctx) {
        ctx.clearRect(0, 0, canvas.width, canvas.height);
        ctx.fillStyle = "#ffffff";
        ctx.fillRect(0, 0, canvas.width, canvas.height);
  
        // 添加随机颜色字符
        for (let i = 0; i < text.length; i++) {
          ctx.font = "30px Arial";
          ctx.fillStyle = `rgb(${Math.floor(Math.random() * 256)}, ${Math.floor(
            Math.random() * 256
          )}, ${Math.floor(Math.random() * 256)})`;
          ctx.fillText(text[i], 10 + i * 20, 30);
        }
  
        // 添加随机噪点
        for (let i = 0; i < 100; i++) {
          ctx.fillStyle = `rgb(${Math.floor(Math.random() * 256)}, ${Math.floor(
            Math.random() * 256
          )}, ${Math.floor(Math.random() * 256)})`;
          ctx.beginPath();
          ctx.arc(
            Math.random() * canvas.width,
            Math.random() * canvas.height,
            1,
            0,
            Math.PI * 2
          );
          ctx.fill();
        }
        // 添加干扰线
        for (let i = 0; i < 5; i++) {
          ctx.strokeStyle = `rgb(${Math.floor(Math.random() * 256)}, ${Math.floor(
            Math.random() * 256
          )}, ${Math.floor(Math.random() * 256)})`;
          ctx.beginPath();
          ctx.moveTo(Math.random() * canvas.width, Math.random() * canvas.height);
          ctx.lineTo(Math.random() * canvas.width, Math.random() * canvas.height);
          ctx.stroke();
        }
      }
    }
    console.log(captchaText)
    return captchaText
  };