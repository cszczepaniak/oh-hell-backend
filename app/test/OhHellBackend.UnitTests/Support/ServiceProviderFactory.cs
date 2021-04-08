using System;
using Microsoft.Extensions.DependencyInjection;
using OhHellBackend.DataAccess.Extensions;

namespace OhHellBackend.UnitTests.Support
{
    public static class ServiceProviderFactory
    {
        public static IServiceProvider Create()
        {
            var services = new ServiceCollection();
            services
                .AddDataAccess()
                .AddDataAccessFakes();
            return services.BuildServiceProvider();
        }
    }
}