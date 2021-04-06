using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;

namespace OhHellBackend.Controllers
{
    [Route("api/[controller]")]
    public class GamesController : ControllerBase
    {
        [HttpPost]
        public async Task<IActionResult> Save()
        {
            return Ok();
        }
    }
}
