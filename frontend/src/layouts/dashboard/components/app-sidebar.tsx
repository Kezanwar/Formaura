import * as React from 'react';
import {
  BookOpen,
  Bot,
  User,
  PieChart,
  Settings2,
  SquareTerminal,
  CreditCard,
  Scroll,
  ScrollText,
  Layers,
  Users
} from 'lucide-react';

import { NavMain } from '@app/layouts/dashboard/components/nav-main';
import NavAccount from '@app/layouts/dashboard/components/nav-account';
import NavUser from '@app/layouts/dashboard/components/nav-user';

import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarHeader,
  SidebarRail
} from '@app/components/ui/sidebar';
import TopSection from './top-section';

const data = {
  navMain: [
    {
      title: 'Forms',
      url: 'forms/manage',
      icon: Layers,
      isActive: true,
      items: [
        {
          title: 'Manage',
          url: 'forms/manage'
        },
        {
          title: 'Create',
          url: 'forms/new'
        }
      ]
    },
    {
      title: 'Affiliates',
      url: 'forms/manage',
      icon: Users,
      isActive: true,
      items: [
        {
          title: 'Manage',
          url: 'forms/manage'
        },
        {
          title: 'Add',
          url: 'forms/create'
        }
      ]
    }
    // {
    //   title: 'Trades',
    //   url: '#',
    //   icon: PieChart,
    //   items: [
    //     {
    //       title: 'General',
    //       url: '#'
    //     },
    //     {
    //       title: 'Team',
    //       url: '#'
    //     },
    //     {
    //       title: 'Billing',
    //       url: '#'
    //     },
    //     {
    //       title: 'Limits',
    //       url: '#'
    //     }
    //   ]
    // },
    // {
    //   title: 'Strategies',
    //   url: '#',
    //   icon: SquareTerminal,
    //   items: [
    //     {
    //       title: 'Genesis',
    //       url: '#'
    //     },
    //     {
    //       title: 'Explorer',
    //       url: '#'
    //     },
    //     {
    //       title: 'Quantum',
    //       url: '#'
    //     }
    //   ]
    // },
    // {
    //   title: 'Reports',
    //   url: '#',
    //   icon: BookOpen,
    //   items: [
    //     {
    //       title: 'Introduction',
    //       url: '#'
    //     },
    //     {
    //       title: 'Get Started',
    //       url: '#'
    //     },
    //     {
    //       title: 'Tutorials',
    //       url: '#'
    //     },
    //     {
    //       title: 'Changelog',
    //       url: '#'
    //     }
    //   ]
    // }
  ],
  account: [
    {
      name: 'Billing',
      url: '#',
      icon: CreditCard
    },
    {
      name: 'Subscription',
      url: '#',
      icon: Settings2
    },
    {
      name: 'Profile',
      url: '#',
      icon: User
    }
  ]
};

export function AppSidebar({ ...props }: React.ComponentProps<typeof Sidebar>) {
  return (
    <Sidebar collapsible="icon" {...props}>
      <SidebarHeader>
        <TopSection />
      </SidebarHeader>
      <SidebarContent>
        <NavMain items={data.navMain} />
        <NavAccount account={data.account} />
      </SidebarContent>
      <SidebarFooter className="mb-2">
        <NavUser />
      </SidebarFooter>
      <SidebarRail />
    </Sidebar>
  );
}
