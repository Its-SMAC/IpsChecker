document.addEventListener("DOMContentLoaded", () => {
  const btnScan = document.getElementById("btnScan");
  const ipInput = document.getElementById("ipInput");
  const resultadosContainer = document.getElementById("resultadosContainer");

  const filtroContainer = document.getElementById("filtroContainer");
  const buscarInput = document.getElementById("buscarInput");

  let listaDeHosts = [];

  const renderizarCartoes = (hosts) => {
    resultadosContainer.innerHTML = "";

    if (hosts.length === 0) {
      resultadosContainer.innerHTML =
        "<p class='text-sm text-gray-500 font-mono col-span-full text-center'>Nenhum resultado corresponde à pesquisa.</p>";
      return;
    }

    hosts.forEach((host) => {
      const statusColor = host.busy
        ? "bg-emerald-50 border-emerald-200 text-emerald-800"
        : "bg-red-50 border-red-200 text-red-800";
      const statusText = host.busy ? "Ativo" : "Inacessível";

      resultadosContainer.innerHTML += `
        <div class="border ${statusColor} rounded-md p-4 shadow-sm flex justify-between items-center font-mono text-sm">
            <div>
                <p class="font-semibold">${host.address}</p>
            </div>
            <span class="text-xs font-semibold uppercase tracking-wider">${statusText}</span>
        </div>
      `;
    });
  };

  btnScan.addEventListener("click", async () => {
    const ipValue = ipInput.value.trim();

    if (!ipValue) {
      alert("Por favor, introduza um endereço IP válido.");
      return;
    }

    btnScan.disabled = true;
    btnScan.innerText = "A verificar...";
    buscarInput.value = "";
    filtroContainer.classList.add("hidden");
    resultadosContainer.innerHTML =
      "<p class='text-sm text-gray-500 font-mono col-span-full text-center'>A executar varrimento na rede...</p>";

    try {
      const response = await fetch("/scan/ip", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ alvo: ipValue }),
      });

      if (!response.ok) throw new Error("Erro na comunicação com o servidor.");

      listaDeHosts = await response.json();

      if (listaDeHosts && listaDeHosts.length > 0) {
        filtroContainer.classList.remove("hidden");
      }

      renderizarCartoes(listaDeHosts);
    } catch (error) {
      resultadosContainer.innerHTML = `<p class='text-sm text-red-600 font-mono col-span-full text-center'>${error.message}</p>`;
    } finally {
      btnScan.disabled = false;
      btnScan.innerText = "Executar Verificação";
    }
  });

  buscarInput.addEventListener("input", () => {
    const termoPesquisa = buscarInput.value.toLowerCase().trim();

    const hostsFiltrados = listaDeHosts.filter((host) => {
      const ipBateCerto = host.address.toLowerCase().includes(termoPesquisa);

      const statusTexto = host.busy ? "ativo" : "inacessível";
      const statusBateCerto = statusTexto.includes(termoPesquisa);

      return ipBateCerto || statusBateCerto;
    });

    renderizarCartoes(hostsFiltrados);
  });
});
