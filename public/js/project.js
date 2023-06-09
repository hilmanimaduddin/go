let dataBlog = [];

function addBlog(event) {
  event.preventDefault();

  let project = document.getElementById("input-project").value;
  let startDate = new Date(document.getElementById("start-date").value);
  let endDate = new Date(document.getElementById("end-date").value);
  let description = document.getElementById("input-description").value;
  let image1 = document.getElementById("upload-image").files;
  let waktuu = new Date();

  // if (project === "") {
  //   return alert("Mohon projeknya diisi dulu ya..");
  // } else if (startDate === "") {
  //   return alert("Mohon tanggalnya diisi dulu ya...");
  // } else if (endDate === "") {
  //   return alert("Mohon tanggalnya diisi dulu ya...");
  // } else if (description === "") {
  //   return alert("Mohon deskripsinya diisi dulu ya...");
  // } else if (image1 == "") {
  //   return alert("Mohon fotonya diupload ya...");
  // }

  image1 = URL.createObjectURL(image1[0]);
  console.log(image1);

  const nodeJs = '<i class="fa-brands fa-node-js" style="color: #000000;"></i>';
  const reactJs = '<i class="fa-brands fa-react" style="color: #000000;"></i>';
  const nextJs =
    '<i class="fa-brands fa-jsfiddle" style="color: #050505;"></i>';
  const typeScript =
    '<i class="fa-brands fa-html5" style="color: #000000;"></i>';

  let Nodejscek = document.getElementById("checkbox1").checked ? nodeJs : "";
  let Nextjscek = document.getElementById("checkbox2").checked ? nextJs : "";
  let Reactjscek = document.getElementById("checkbox3").checked ? reactJs : "";
  let Typescriptcek = document.getElementById("checkbox4").checked
    ? typeScript
    : "";

  console.log(Nodejscek, Nextjscek, Reactjscek, Typescriptcek);

  let waktu = Math.abs(endDate - startDate);
  let bulan = Math.floor(waktu / (1000 * 60 * 60 * 24 * 30));
  let hari = Math.floor(waktu / (1000 * 60 * 60 * 24));

  let cobaa = {
    project,
    startDate,
    endDate,
    description,
    image1,
    Nodejscek,
    Nextjscek,
    Reactjscek,
    Typescriptcek,
    bulan,
    hari,
    waktuu,
  };

  console.log(cobaa);

  dataBlog.push(cobaa);
  console.log(dataBlog);

  rendercobaa();
}

function rendercobaa() {
  document.getElementById("blog-under").innerHTML = "";

  for (let index = 0; index < dataBlog.length; index++) {
    document.getElementById("blog-under").innerHTML += `
    <div class="card mycard rounded border m-3" style="width: 18rem">
        <img
          src="${dataBlog[index].image1}"
          alt="image"
          class="object-fit-cover p-2"
          style="height: 200px"
        />
        <h3 class="ms-2 me-2">${dataBlog[index].project}</h3>
        <p class="text-secondary ms-2 fs-6">
          Durasi : ${dataBlog[index].bulan} Bulan ${dataBlog[index].hari} Hari
        </p>
        <p class="ms-2 me-2">${dataBlog[index].description}</p>
        <div class="ms-2">
          ${dataBlog[index].Nodejscek} ${dataBlog[index].Nextjscek}
          ${dataBlog[index].Reactjscek} ${dataBlog[index].Typescriptcek}
        </div>
        <div>
          <p style="font-size: 15px; color: grey; float: right; margin-right: 15px;">
            ${getDistanceTime(dataBlog[index].waktuu)}
          </p>
          </div>
          <div class="mt-0 ms-1 me-1 row">
          <button class="btn btn-dark col m-2">Edit</button>
          <button class="btn btn-dark col m-2">Delete</button>
          </div>
          </div>
          `;
  }
}

function getDistanceTime(time) {
  let timeNow = new Date();
  let timePost = time;

  let distance = timeNow - timePost;
  console.log(distance);

  let distanceDay = Math.floor(distance / (1000 * 60 * 60 * 24));
  let distanceHour = Math.floor(distance / (1000 * 60 * 60));
  let distanceMinute = Math.floor(distance / (1000 * 60));
  let distanceSecond = Math.floor(distance / 1000);

  if (distanceDay > 0) {
    return `${distanceDay} Day ago`;
  } else if (distanceHour > 0) {
    return `${distanceHour} Hour ago`;
  } else if (distanceMinute > 0) {
    return `${distanceMinute} Minute ago`;
  } else {
    return `${distanceSecond} Second ago`;
  }
}

setInterval(function () {
  rendercobaa();
}, 5000);
