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
    public sealed class ProjectScheduledTriggerDaysPerMonthSchedule
    {
        /// <summary>
        /// Which date of the month to run the trigger. String number between 1 - 31 Incl. or L for the last day of the month.
        /// </summary>
        public readonly string? DateOfMonth;
        /// <summary>
        /// Which ordinal day of the week to run the trigger on. String number between 1 - 4 Incl. or L for the last occurrence of day*of*week for the month.
        /// </summary>
        public readonly string? DayNumberOfMonth;
        /// <summary>
        /// Which day of the week to run the trigger on. Required when monthly*schedule*type is set to 'DayOfMonth'.
        /// </summary>
        public readonly string? DayOfWeek;
        /// <summary>
        /// The type of monthly schedule to run the trigger
        /// </summary>
        public readonly string MonthlyScheduleType;
        /// <summary>
        /// The time of day to start the trigger.
        /// </summary>
        public readonly string StartTime;

        [OutputConstructor]
        private ProjectScheduledTriggerDaysPerMonthSchedule(
            string? dateOfMonth,

            string? dayNumberOfMonth,

            string? dayOfWeek,

            string monthlyScheduleType,

            string startTime)
        {
            DateOfMonth = dateOfMonth;
            DayNumberOfMonth = dayNumberOfMonth;
            DayOfWeek = dayOfWeek;
            MonthlyScheduleType = monthlyScheduleType;
            StartTime = startTime;
        }
    }
}
