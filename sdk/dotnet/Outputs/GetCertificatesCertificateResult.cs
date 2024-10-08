// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy.Outputs
{

    [OutputType]
    public sealed class GetCertificatesCertificateResult
    {
        public readonly string Archived;
        /// <summary>
        /// The encoded data of the certificate.
        /// </summary>
        public readonly string CertificateData;
        /// <summary>
        /// Specifies the archive file format used for storing cryptography objects in the certificate. Valid formats are `Der`, `Pem`, `Pkcs12`, or `Unknown`.
        /// </summary>
        public readonly string CertificateDataFormat;
        /// <summary>
        /// A list of environment IDs associated with this resource.
        /// </summary>
        public readonly ImmutableArray<string> Environments;
        /// <summary>
        /// Indicates if the certificate has a private key.
        /// </summary>
        public readonly bool HasPrivateKey;
        /// <summary>
        /// The unique ID for this resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Indicates if the certificate has expired.
        /// </summary>
        public readonly bool IsExpired;
        public readonly string IssuerCommonName;
        public readonly string IssuerDistinguishedName;
        public readonly string IssuerOrganization;
        /// <summary>
        /// The name of this resource.
        /// </summary>
        public readonly string Name;
        public readonly string NotAfter;
        public readonly string NotBefore;
        public readonly string Notes;
        /// <summary>
        /// The password associated with this resource.
        /// </summary>
        public readonly string Password;
        public readonly string ReplacedBy;
        public readonly bool SelfSigned;
        public readonly string SerialNumber;
        public readonly string SignatureAlgorithmName;
        public readonly string SpaceId;
        public readonly ImmutableArray<string> SubjectAlternativeNames;
        public readonly string SubjectCommonName;
        public readonly string SubjectDistinguishedName;
        public readonly string SubjectOrganization;
        /// <summary>
        /// A list of tenant tags associated with this resource.
        /// </summary>
        public readonly ImmutableArray<string> TenantTags;
        /// <summary>
        /// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
        /// </summary>
        public readonly string TenantedDeploymentParticipation;
        /// <summary>
        /// A list of tenant IDs associated with this resource.
        /// </summary>
        public readonly ImmutableArray<string> Tenants;
        public readonly string Thumbprint;
        public readonly int Version;

        [OutputConstructor]
        private GetCertificatesCertificateResult(
            string archived,

            string certificateData,

            string certificateDataFormat,

            ImmutableArray<string> environments,

            bool hasPrivateKey,

            string id,

            bool isExpired,

            string issuerCommonName,

            string issuerDistinguishedName,

            string issuerOrganization,

            string name,

            string notAfter,

            string notBefore,

            string notes,

            string password,

            string replacedBy,

            bool selfSigned,

            string serialNumber,

            string signatureAlgorithmName,

            string spaceId,

            ImmutableArray<string> subjectAlternativeNames,

            string subjectCommonName,

            string subjectDistinguishedName,

            string subjectOrganization,

            ImmutableArray<string> tenantTags,

            string tenantedDeploymentParticipation,

            ImmutableArray<string> tenants,

            string thumbprint,

            int version)
        {
            Archived = archived;
            CertificateData = certificateData;
            CertificateDataFormat = certificateDataFormat;
            Environments = environments;
            HasPrivateKey = hasPrivateKey;
            Id = id;
            IsExpired = isExpired;
            IssuerCommonName = issuerCommonName;
            IssuerDistinguishedName = issuerDistinguishedName;
            IssuerOrganization = issuerOrganization;
            Name = name;
            NotAfter = notAfter;
            NotBefore = notBefore;
            Notes = notes;
            Password = password;
            ReplacedBy = replacedBy;
            SelfSigned = selfSigned;
            SerialNumber = serialNumber;
            SignatureAlgorithmName = signatureAlgorithmName;
            SpaceId = spaceId;
            SubjectAlternativeNames = subjectAlternativeNames;
            SubjectCommonName = subjectCommonName;
            SubjectDistinguishedName = subjectDistinguishedName;
            SubjectOrganization = subjectOrganization;
            TenantTags = tenantTags;
            TenantedDeploymentParticipation = tenantedDeploymentParticipation;
            Tenants = tenants;
            Thumbprint = thumbprint;
            Version = version;
        }
    }
}
